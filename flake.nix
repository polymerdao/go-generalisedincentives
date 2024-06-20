{
  description = "Flake for go-generalisedincentives";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    foundry.url = "github:shazow/foundry.nix/monthly"; # Use monthly branch for permanent releases
  };

  outputs = { self, nixpkgs, flake-utils, foundry }:
    (flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = import nixpkgs {
            inherit system;
            overlays = [ foundry.overlay ];
          };
        in
        {
          devShells.default = if system == "aarch64-darwin" then
            import ./shell-darwin.nix { inherit pkgs; }
          else if system == "x86_64-linux" then
            import ./shell-linux.nix { inherit pkgs; }
          else
            throw "${system} is not supported by this flake.";
        })
    );
}
