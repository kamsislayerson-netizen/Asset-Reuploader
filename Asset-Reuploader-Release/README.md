# Asset Reuploader - Release Package v1.3.1

Complete package for Asset Reuploader with all required files.

## Contents

### `AssetReuploader.exe` (Windows Executable)
- **Version:** 1.3.1
- **Size:** ~7.0 MB
- **Compatible Plugin:** v1.3.1
- **Usage:** Run this on Windows to start the Asset Reuploader server
- **Command:** `AssetReuploader.exe` (runs with default config.ini settings)

### `AssetReuploader-linux` (Linux Executable)
- **Version:** 1.3.1
- **Size:** ~6.7 MB
- **Compatible Plugin:** v1.3.1
- **Usage:** Run this on Linux to start the Asset Reuploader server
- **Command:** `./AssetReuploader-linux` (runs with default config.ini settings)

### `AssetReuploader1.3.1.rbxm` (Roblox Plugin)
- **Version:** 1.3.1
- **Size:** ~48 KB
- **Format:** Roblox Model File (.rbxm)
- **Usage:** Import this plugin into Roblox Studio
- **Installation:**
  1. Open Roblox Studio
  2. Go to "Plugins" → "Manage Plugins" → "Create New Plugin"
  3. Or drag-and-drop this file into your Plugins folder in Roblox

### `config.ini` (Configuration File)
- Default configuration file for the Asset Reuploader executable
- Contains server settings like port, cookie file location, etc.
- Edit this file to customize:
  - Server port (default: typically 3000)
  - Cookie storage location
  - Other runtime settings

## Quick Start

### Windows
1. Extract all files to a folder
2. (Optional) Edit `config.ini` for custom settings
3. Run `AssetReuploader.exe`
4. In Roblox Studio, install the `AssetReuploader1.3.1.rbxm` plugin
5. Use the plugin to reupload assets

### Linux
1. Extract all files to a folder
2. (Optional) Edit `config.ini` for custom settings
3. Run `./AssetReuploader-linux`
4. In Roblox Studio, install the `AssetReuploader1.3.1.rbxm` plugin
5. Use the plugin to reupload assets

## Features

✅ **Animations** - Reupload animation assets for free  
✅ **Audio/Sound** - Reupload sound assets for free  
✅ **Meshes** - Reupload mesh assets for free  
✅ **Images** - Reupload image assets for free  
✅ **Rate Limiting** - Slow, steady uploads to avoid rate limits  
✅ **Auto-Retry** - Automatic retry on temporary failures  
✅ **Asset Tracking** - Tracks which assets have been reuploaded

## Requirements

- **Windows:** Windows 7 or later (for .exe file)
- **Linux:** Linux with x86_64 architecture (for Linux binary)
- **Roblox Studio:** Latest version
- **Roblox Account:** To authenticate and reupload assets

## Configuration

Edit `config.ini` to customize:
- Server port (ensure firewall allows access)
- Cookie storage location (where Roblox cookies are saved)
- Other application settings

## Support & Community

- **Discord:** https://discord.gg/XTEtUqPTat
- **GitHub Issues:** https://github.com/kamsislayerson-netizen/Asset-Reuploader/issues
- **Repository:** https://github.com/kamsislayerson-netizen/Asset-Reuploader

## License

Licensed under GPL-3.0. See LICENSE file in the repository for details.

---

**Built:** August 12, 2026  
**Version:** 1.3.1 + Image Support  
**Status:** Ready to use
