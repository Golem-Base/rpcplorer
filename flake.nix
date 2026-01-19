{
  description = "RPCPlorer";
  inputs = {

    nixpkgs = {
      url = "github:NixOS/nixpkgs/nixos-unstable";
    };

    systems.url = "github:nix-systems/default";

    templ = {
      url = "github:a-h/templ/v0.3.833";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        nixpkgs-unstable.follows = "nixpkgs";
      };
    };

    gorefresh = {
      url = "github:draganm/gorefresh/v0.0.4";
      inputs = {
        nixpkgs.follows = "nixpkgs";
      };
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      systems,
      templ,
      gorefresh,
      ...
    }@inputs:
    let
      eachSystem = nixpkgs.lib.genAttrs (import systems) (
        system:
        let
          pkgs = import nixpkgs {
            inherit system;
            config = {
              allowUnfree = true;
            };
          };
        in
        {
          default = pkgs.callPackage ./package.nix { };
          devShell = pkgs.mkShell {
            shellHook = ''
              # Set here the env vars you want to be available in the shell
            '';
            hardeningDisable = [ "all" ];

            packages = with pkgs; [
              go
              shellcheck
              (templ.packages.${system}.default)
              (gorefresh.packages.${system}.default)
              overmind
            ];
          };
        }
      );

    in
    {
      devShells = eachSystem;
      packages = eachSystem;
    };
}
