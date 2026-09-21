{
  description = "gocard - A Go library for representing and manipulating playing cards and decks.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable-small";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
    nur.url = "github:nix-community/NUR";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    gomod2nix,
    nur,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [gomod2nix.overlays.default nur.overlays.default];
          config.allowUnfreePredicate = pkg: builtins.elem (pkgs.lib.getName pkg) ["goreleaser-pro"];
        };
        lib = pkgs.lib;
        version =
          if (self ? shortRev)
          then self.shortRev
          else "dev";
      in {
        packages.default = pkgs.buildGoApplication {
          pname = "gocard";
          inherit version;

          src = lib.fileset.toSource {
            root = ./.;
            fileset = lib.fileset.unions [
              ./go.mod
              ./go.sum
              ./gomod2nix.toml
              (lib.fileset.fileFilter (file: lib.hasSuffix ".go" file.name) ./.)
            ];
          };

          modules = ./gomod2nix.toml;

          subPackages = ["."];

          ldflags = [
            "-s"
            "-w"
            "-X github.com/asphaltbuffet/gocard/internal/version.Version=${version}"
          ];

          meta = with lib; {
            description = "A Go library for representing and manipulating playing cards and decks.";
            homepage = "https://github.com/asphaltbuffet/gocard";
            license = licenses.mit;
            mainProgram = "gocard";
          };
        };

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
            pkgs.nur.repos.goreleaser.goreleaser-pro
            gomod2nix.packages.${system}.default
            gh
          ];

          shellHook = ''
            mise trust --all
          '';
        };
      }
    )
    // {
      overlays.default = final: prev: {
        gocard = self.packages.${prev.system}.default;
      };
    };
}
