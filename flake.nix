{
  description = "gocard - A Go library for representing and manipulating playing cards and decks.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable-small";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {inherit system;};
      in {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            jujutsu
            jjui
            mise
            ripgrep
            fd
            sd
            gopls
            nixd
            gh
          ];

          shellHook = ''
            mise trust --all
          '';
        };
      }
    );
}
