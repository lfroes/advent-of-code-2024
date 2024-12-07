local function readFile(filename)
	local file = io.open(filename, "r")
	if not file then
		error("Not possible to open the file" .. filename)
	end
	local data = file:read("*all")
	file:close()
	return data
end

local function filterLines(lines)
	print("Lines:", table.concat(lines, " | "))
	local filtered = {}
	for _, line in ipairs(lines) do
		local trimmed = line:match("^%s*(.-)%s*$")
		if trimmed ~= "" then
			table.insert(filtered, trimmed)
		end
	end

	return filtered
end

local function split(input, delimiter)
	local result = {}
	for match in (input .. delimiter):gmatch("(.-)" .. delimiter) do
		table.insert(result, match)
	end
	return result
end

local function validateOrder(order, deps)
	for j = 1, #order do
		for _, dep in ipairs(deps) do
			local num1, num2 = dep:match("([^|]+)|([^|]+)")
			if num2 == order[j] then
				local foundIndex = nil
				for k = 1, #order do
					if order[k] == num1 then
						foundIndex = k
						break
					end
				end

				if foundIndex and foundIndex > j then
					return false
				end
			end
		end
	end
	return true
end

local function adjustOrder(order, deps, convertedValidTops)
	if validateOrder(order, deps) then
		table.insert(convertedValidTops, table.concat(order, ","))
		return order
	end

	for i = 1, #order do
		for _, dep in ipairs(deps) do
			local num1, num2 = dep:match("([^|]+)|([^|]+)")
			local idx1, idx2
			for idx, num in ipairs(order) do
				if num == num1 then
					idx1 = idx
				end
				if num == num2 then
					idx2 = idx
				end
			end
			if idx1 and idx2 and idx2 < idx1 then
				order[idx1], order[idx2] = order[idx2], order[idx1]
			end
		end
	end

	return adjustOrder(order, deps, convertedValidTops)
end

local data = readFile("input.txt")
local depsRaw, topRaws = data:match("^(.-)\n\n(.-)$")

local deps = filterLines(split(depsRaw, "\n"))
local tops = filterLines(split(topRaws, "\n"))
local convertedValidTops = {}

for _, top in ipairs(tops) do
	local orderFiltered = {}
	for num in top:gmatch("[^,]+") do
		table.insert(orderFiltered, num)
	end
	if not validateOrder(orderFiltered, deps) then
		adjustOrder(orderFiltered, deps, convertedValidTops)
	end
end

local acc = 0
for _, top in ipairs(convertedValidTops) do
	local nums = {}
	for num in top:gmatch("[^,]+") do
		table.insert(nums, num)
	end
	local middle = nums[math.floor(#nums / 2) + 1]
	acc = acc + tonumber(middle)
end

print("Converted:", table.concat(convertedValidTops, " | "))
print("ValidTops:", acc)
