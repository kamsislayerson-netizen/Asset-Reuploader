--!strict
local getFilterOptions = require("../GetFilterOptions")
local reuploadIds = require("../ReuploadIds")
local UiLibrary = require("../../UiLibrary")

local Selection = game:GetService("Selection")

return function(ui: UiLibrary.Ui, plugin: Plugin)
	local tab = ui:CreateTab("Image")
	
	tab:CreateButton("Reupload", function()
		local filter = getFilterOptions(plugin, game:GetDescendants())
		table.insert(filter.WhitelistedInstances, "StringValue")
		reuploadIds(plugin, ui, filter, "Image", _G.getPlaceList()) -- forgive me
	end)
	
	tab:CreateButton("Reupload Selected", function()
		local filter = {
			WhitelistedInstances = { "StringValue", "NumberValue", "IntValue", "LuaSourceContainer" },
			Instances = Selection:Get()
		}

		reuploadIds(plugin, ui, filter, "Image", _G.getPlaceList()) -- forgive me
	end)
end
