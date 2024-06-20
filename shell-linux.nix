{ pkgs }:

pkgs.mkShell {
  packages = [
    pkgs.go
    pkgs.foundry-bin
    pkgs.go-ethereum
  ];
}
