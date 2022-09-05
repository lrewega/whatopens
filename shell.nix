{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    git
    go_1_19
    gopls 
    gotools 
    darwin.apple_sdk_11_0.MacOSX-SDK
    darwin.apple_sdk_11_0.frameworks.AppKit
    darwin.apple_sdk_11_0.frameworks.Foundation
  ];
}
