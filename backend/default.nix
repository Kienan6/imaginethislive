with import <nixpkgs> {};

stdenv.mkDerivation {

  name = "authapi-server";

  buildInputs = with pkgs; [
    gnumake
    go_1_19
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$PATH:$HOME/go/bin
  '';

}
