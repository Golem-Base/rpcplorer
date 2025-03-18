{
  description = "RPCPlorer";
  inputs = {

    nixpkgs = {
      url = "github:NixOS/nixpkgs/master";
    };

    systems.url = "github:nix-systems/default";

    templ.url = "github:a-h/templ/v0.3.833";

  };

  outputs =
    {
      self,
      nixpkgs,
      systems,
      templ,
      ...
    }@inputs:
    let
      eachSystem =
        nixpkgs.lib.genAttrs (import systems) (
          system:
          let pkgs = import nixpkgs {
            inherit system;
            config = {
              allowUnfree = true;
            };
          };
          in
          {
            default = pkgs.mkShell {
              shellHook = ''
                # Set here the env vars you want to be available in the shell
              '';
              hardeningDisable = [ "all" ];

              packages = with pkgs; [
                go
                shellcheck
                (templ.packages.${system}.default)
              ];
            };
          }
        );

    in
    {
      devShells = eachSystem;
    };
}
