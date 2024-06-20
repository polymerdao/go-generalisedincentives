{ pkgs }:

pkgs.mkShell {
  buildInputs = [
    pkgs.darwin.apple_sdk.frameworks.Security
    pkgs.pkg-config
    pkgs.openssl
  ];

  packages = [
    pkgs.go
    pkgs.foundry-bin
    pkgs.go-ethereum
  ];

  shellHook = ''
  export CPATH="${pkgs.darwin.Libsystem}/include";
  '';         
}
