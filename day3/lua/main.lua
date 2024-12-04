local read_memory = {}

function read_memory:main(path)
	local file = io.open(path, "r")

	if not file then
		error("Could not open file " .. path)
	end

	local patterns = {
		{ type = "mul", pattern = "mul%(%d+,%d+%)" },
		{ type = "do", pattern = "do%(%w*%)" },
		{ type = "dont", pattern = "don't%(%w*%)" },
	}

	local results = {}
	local current_index = 0

	for line in file:lines() do
		local line_length = #line
		local line_start_index = current_index

		for _, pat in ipairs(patterns) do
			for match_start, match_end in line:gmatch("()" .. pat.pattern .. "()") do
				local absolute_start = line_start_index + match_start
				table.insert(results, {
					type = pat.type,
					value = line:sub(match_start, match_end - 1),
					index = absolute_start,
				})
			end
		end

		current_index = current_index + line_length + 1
	end

	table.sort(results, function(a, b)
		return a.index < b.index
	end)

	local solution = 0
	local is_enabled = true

	for _, result in ipairs(results) do
		if result.type == "do" then
			is_enabled = true
		elseif result.type == "dont" then
			is_enabled = false
		elseif result.type == "mul" and is_enabled then
			local n1, n2 = result.value:match("mul%((%d+),(%d+)%)")
			solution = solution + (tonumber(n1) * tonumber(n2))
		end
	end

	print("Final Solution:", solution)
	file:close()
end

read_memory:main("input.txt")
