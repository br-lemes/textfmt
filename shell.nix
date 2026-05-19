with import <nixpkgs> { };

mkShell {
  nativeBuildInputs = [
    gh
    git
    gnumake
    go
    openssh
  ];
}
