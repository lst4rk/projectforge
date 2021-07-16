source = ["./build/dist/darwin_amd64_darwin_amd64/projectforge"]
bundle_id = "com.kyleu.projectforge"

//notarize {
//  path = "./build/dist/projectforge_desktop_0.0.0_macos_x86_64.dmg"
//  bundle_id = "com.kyleu.projectforge"
//}

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "C6S478FYLD"
}

dmg {
  output_path = "./build/dist/projectforge_0.0.0_macos_x86_64.dmg"
  volume_name = "Project Forge"
}

zip {
  output_path = "./build/dist/projectforge_0.0.0_macos_x86_64_notarized.zip"
}
