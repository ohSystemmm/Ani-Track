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
      pkgs = nixpkgs.lib.mkShell {
        inherit (nixpkgs) lib;
        go = nixpkgs.go_1_20;
      };
    in
    {
      defaultPackage = pkgs.stdenv.mkDerivation {
        pname = "ani-track";
        version = "0.5.0";

        src = ./.;

        buildInputs = [ pkgs.go ];

        buildPhase = ''
          go build -o ${pname}
        '';

        installPhase = ''
          mkdir -p $out/bin
          mv ${pname} $out/bin/
        '';
      };
    };

    devShell = pkgs.mkShell {
      buildInputs = [ pkgs.go ];
    };
  };
}

