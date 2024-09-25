#   __ _       _                _
# / _| | __ _| | _____   _ __ (_)_  __
# | |_| |/ _` | |/ / _ \ | '_ \| \ \/ /
# |  _| | (_| |   <  __/_| | | | |>  <
# |_| |_|\__,_|_|\_\___(_)_| |_|_/_/\_\
# 
# by ohSystemmm <3

{
  description = "A CLI based Anime Tracker";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }: {
    packages = let
      pkgs = import nixpkgs {
        inherit (nixpkgs) system;
      };
    in
    {
      defaultPackage = pkgs.stdenv.mkDerivation {
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
    };

    devShell = pkgs.mkShell {
      buildInputs = [ pkgs.go_1_20 ];
    };
  };
}
