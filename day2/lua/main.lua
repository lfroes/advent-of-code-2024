local solution = {}
local safe_lists = 0

function Is_safe(numbers)
	local is_crescent = true
	local is_decrescent = true

	-- Iterate over numbers

	for i = 1, #numbers - 1 do
		local diff = math.abs(numbers[i + 1] - numbers[i])

		if diff > 3 or diff < 1 then
			return false
		end

		if numbers[i + 1] < numbers[i] then
			is_crescent = false
		end

		if numbers[i + 1] > numbers[i] then
			is_decrescent = false
		end
	end

	return is_crescent or is_decrescent
end

function solution:main(input)
	local file = io.open(input, "r")

	if not file then
		error("Could not open file " .. input)
	end

	for line in file:lines() do
		local numbers = {}

		for num in string.gmatch(line, "%d+") do
			table.insert(numbers, tonumber(num))
		end

		if Is_safe(numbers) then
			safe_lists = safe_lists + 1
		else
			for i = 1, #numbers do
				local temp = { table.unpack(numbers) }
				table.remove(temp, i)
				if Is_safe(temp) then
					safe_lists = safe_lists + 1
					break
				end
			end
		end
	end

	file:close()
	print("Safe Lists:", safe_lists)
end

solution:main("./input.txt")
