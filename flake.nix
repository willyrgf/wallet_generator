{
  description = "Go project with buildGoModule and vendor hash";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        wallet_generator = pkgs.buildGoModule {
          pname = "wallet_generator";
	  version = "0.0.1";
          src = ./.;
          #vendorHash = pkgs.lib.fakeHash;
	  vendorHash = "sha256-dp/VqglYmWMEmK5NXdUP6ym24A26kSZYciYDL3vh/AU=";
        };
      in
      {
        defaultPackage = wallet_generator;
        apps.wallet_generator = {
          type = "app";
          program = "${wallet_generator}/bin/wallet_generator";
        };
      }
    );
}
