{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gopls
  ];

  GOROOT = "${pkgs.go}/share/go";
}
