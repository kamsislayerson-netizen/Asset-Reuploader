# 🚀 Quick Start Guide - Asset Reuploader v1.3.1

## 5-Minute Setup

### Step 1: Prepare the Executable
**Windows:**
```bash
# Just run the exe
AssetReuploader.exe
```

**Linux:**
```bash
# Make it executable (if needed)
chmod +x AssetReuploader-linux

# Run it
./AssetReuploader-linux
```

### Step 2: Install the Plugin
1. **In Roblox Studio:**
   - Go to `Plugins` menu → `Import File` (or `Manage Plugins`)
   - Select `AssetReuploader1.3.1.rbxm`
   - Click Import

2. **Verify Installation:**
   - You should see a new "Asset Reuploader" tab in the top menu
   - Tab options: Animation, Sound, Mesh, Image, Replace, Filter, Settings

### Step 3: Configure Settings (Optional)
Edit `config.ini` if you need to:
- Change the server port
- Set custom cookie location
- Adjust other settings

Default port is typically **3000**.

### Step 4: Reupload Assets

1. **In Roblox Studio:**
   - Open your game
   - Click the Asset Reuploader tab
   - Select asset type: **Animation**, **Sound**, **Mesh**, or **Image**
   - Click **"Reupload"** to process all assets
   - Or click **"Reupload Selected"** for specific instances

2. **Monitor Progress:**
   - Check the output console for progress
   - Each asset shows: `[count/total] Asset Name (ID)`

3. **Asset Mapping:**
   - Old IDs are replaced with new ones automatically
   - In the Replace tab, you can manually adjust mappings if needed

## Asset Type Options

| Asset Type | File Extension | Supported Locations |
|-----------|-----------------|-------------------|
| **Animation** | N/A | Animation instances |
| **Sound** | .ogg, .mp3 | Sound instances |
| **Mesh** | N/A | MeshPart, SpecialMesh, CharacterMesh |
| **Image** | .png, .jpg, .bmp | String values with rbxassetid:// |

## Common Issues

### ❌ "Plugin is busy"
- Wait for the current reupload to finish
- Check console for errors

### ❌ "Unable to connect to localhost"
- Make sure `AssetReuploader.exe` is running
- Check if the port is correct (default: 3000)
- Firewall may be blocking - allow access

### ❌ "Plugin needs script injection permission"
- Right-click Roblox Studio shortcut
- Select "Run as administrator"
- Or enable script editing in settings

### ❌ "Authentication required to access asset"
- Click "Reupload" again - it will prompt for re-authentication
- Your Roblox session may have expired

## Tips & Best Practices

✅ **Always backup first** - Save your game before reuploading  
✅ **Do small batches first** - Test with 5-10 assets before doing all  
✅ **Keep the exe running** - Don't close it while the plugin is active  
✅ **Monitor the console** - Watch for any error messages  
✅ **Be patient** - Uploading is rate-limited for stability  

## Filter Options

Before reuploading, configure filters:
1. Click the **Filter** tab
2. Choose which instances to search:
   - Scripts (Scripts, LocalScripts, ModuleScripts)
   - String/Number values (for manual ID storage)
   - Specific asset types (MeshPart, SpecialMesh, etc.)
3. Click **Reupload** to process filtered results

## Settings

Click the **Settings** tab to:
- Export results to JSON
- Set port number
- Configure other options

## Need Help?

📖 See `README.md` for full documentation  
💬 Join Discord: https://discord.gg/XTEtUqPTat  
🐛 Report issues: https://github.com/kamsislayerson-netizen/Asset-Reuploader/issues  

---

**You're all set!** Happy reuploading! 🎉
