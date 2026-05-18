with import <nixpkgs> { };

mkShell {
  nativeBuildInputs = [
    git
    gnumake
    go
    pkgsCross.mingwW64.stdenv.cc
  ];
}
