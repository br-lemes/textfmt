with import <nixpkgs> { };

mkShell {
  nativeBuildInputs = [
    git
    gnumake
    go
  ];
}
