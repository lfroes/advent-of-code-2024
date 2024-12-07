local function readFile(filename)
	local file = io.open(filename, "r")
	if not file then
		error("Erro ao abrir o arquivo: " .. filename)
	end
	local data = file:read("*all")
	file:close()
	return data
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
			num1 = num1 and num1:match("^%s*(.-)%s*$")
			num2 = num2 and num2:match("^%s*(.-)%s*$")

			if num2 == order[j] then
				local foundIndex = nil
				for k = 1, #order do
					if order[k] == num1 then
						foundIndex = k
						break
					end
				end

				if foundIndex then
					if foundIndex > j then
						return false
					end
				end
			end
		end
	end
	return true
end

local data = readFile("input.txt")
local parts = split(data, "\n\n")
if #parts < 2 then
	error("Formato inválido do arquivo")
end

local deps = split(parts[1], "\n")
local tops = split(parts[2], "\n")
local validTops = {}

for _, top in ipairs(tops) do
	local orderFiltered = split(top, ",")
	local viable = validateOrder(orderFiltered, deps)

	if viable then
		table.insert(validTops, top)
	end
end

local acc = 0
for _, top in ipairs(validTops) do
	local nums = split(top, ",")
	local middle = nums[math.floor(#nums / 2) + 1]
	local middleValue = tonumber(middle)
	if middleValue then
		acc = acc + middleValue
	else
		print("Erro ao converter número:", middle)
	end
end

print("ValidTops:", acc)
