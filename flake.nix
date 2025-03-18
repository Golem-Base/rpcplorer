{
  description = "RPCPlorer";
  inputs = {

    nixpkgs = {
      url = "github:NixOS/nixpkgs/master";
    };

    systems.url = "github:nix-systems/default";

  };

  outputs =
    {
      self,
      nixpkgs,
      systems,
      ...
    }@inputs:
    let
      eachSystem =
        f:
        nixpkgs.lib.genAttrs (import systems) (
          system:
          f (
            import nixpkgs {
              inherit system;
              config = {
                allowUnfree = true;
              };
            }
          )
        );

    in
    {
  

      devShells = eachSystem (pkgs: {
        default = pkgs.mkShell {
          shellHook = ''
            # Set here the env vars you want to be available in the shell
          '';
          hardeningDisable = [ "all" ];

          packages = with pkgs; [
            go
            shellcheck
          ];
        };
      });
    };
}
