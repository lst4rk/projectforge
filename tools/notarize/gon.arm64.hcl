source = ["./build/dist/darwin_arm64_darwin_arm64/projectforge"]
bundle_id = "com.kyleu.projectforge"

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "C6S478FYLD"
}

dmg {
  output_path = "./build/dist/projectforge_0.0.0_macos_arm64.dmg"
  volume_name = "Project Forge"
}

zip {
  output_path = "./build/dist/projectforge_0.0.0_macos_arm64_notarized.zip"
}
