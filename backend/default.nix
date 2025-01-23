with import <nixpkgs> {};

stdenv.mkDerivation {

  name = "itl-backend";

  buildInputs = with pkgs; [
    gnumake
    go_1_23
    postgresql_14
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$PATH:$HOME/go/bin
  '';

}
