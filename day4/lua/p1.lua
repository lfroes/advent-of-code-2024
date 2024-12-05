local letterValues = { X = 1, M = 2, A = 3, S = 4 }

local function calculateHash(word)
	local hash = 0
	for i = 1, #word do
		local char = word:sub(i, i)
		hash = hash * 21 + (letterValues[char] or 0)
	end
	return hash
end

local function isOnBounds(row, col, rows, cols)
	return row >= 1 and row <= rows and col >= 1 and col <= cols
end

local function findOccurrences(matrix, targetWord, targetHash)
	local rows = #matrix
	local cols = #matrix[1]
	local wordLength = #targetWord
	local occurrences = 0

	local directions = {
		{ 0, 1 },
		{ 1, 0 },
		{ 1, 1 },
		{ 1, -1 },
		{ -1, 0 },
		{ 0, -1 },
		{ -1, 1 },
		{ -1, -1 },
	}

	for row = 1, rows do
		for col = 1, cols do
			for _, dir in ipairs(directions) do
				local dRow, dCol = dir[1], dir[2]
				local hash = 0
				local valid = true

				for i = 0, wordLength - 1 do
					local newRow = row + dRow * i
					local newCol = col + dCol * i

					if not isOnBounds(newRow, newCol, rows, cols) then
						valid = false
						break
					end

					local char = matrix[newRow][newCol]
					hash = hash * 21 + (letterValues[char] or 0)
				end

				if valid and hash == targetHash then
					occurrences = occurrences + 1
				end
			end
		end
	end

	return occurrences
end

local function readMatrix(fileName)
	local matrix = {}
	for line in io.lines(fileName) do
		local row = {}
		for char in line:gmatch(".") do
			table.insert(row, char)
		end
		table.insert(matrix, row)
	end
	return matrix
end

local inputFileName = "input.txt"
local targetWord = "XMAS"
local targetHash = calculateHash(targetWord)

local matrix = readMatrix(inputFileName)
local result = findOccurrences(matrix, targetWord, targetHash)

print("The result is", result)
