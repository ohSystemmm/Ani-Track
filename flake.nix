#   __ _       _                _
#  / _| | __ _| | _____   _ __ (_)_  __
# | |_| |/ _` | |/ / _ \ | '_ \| \ \/ /
# |  _| | (_| |   <  __/_| | | | |>  <
# |_| |_|\__,_|_|\_\___(_)_| |_|_/_/\_\
# 
# Maintainer: ohSystemmm
# Contributor: Y2kun

{
  description = "A simple CLI tool for tracking anime and manga progress.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }: let
    pkgs = import nixpkgs {
      system = builtins.currentSystem;
    };
  in {
    packages.${pkgs.system} = pkgs.stdenv.mkDerivation {
      pname = "ani-track";
      version = "0.5.0";

      src = ./.;

      buildInputs = [ pkgs.go_1_20 ];

      buildPhase = ''
        go build -o ani-track
      '';

      installPhase = ''
        mkdir -p $out/bin
        mv ani-track $out/bin/
      '';
    };

    devShell.${pkgs.system} = pkgs.mkShell {
      buildInputs = [ pkgs.go_1_20 ];
    };
  };
}