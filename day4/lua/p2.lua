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

local function isOnBounds(row, col, rows, cols)
	return row >= 1 and row <= rows and col >= 1 and col <= cols
end

local function findOccurrences(matrix)
	local rows = #matrix
	local cols = #matrix[1]
	local countA = 0
	local occurrences = 0

	for row = 1, rows do
		for col = 1, cols do
			if matrix[row][col] ~= "A" then
				goto continue
			end
			countA = countA + 1

			if
				not isOnBounds(row - 1, col - 1, rows, cols)
				or not isOnBounds(row + 1, col + 1, rows, cols)
				or not isOnBounds(row - 1, col + 1, rows, cols)
				or not isOnBounds(row + 1, col - 1, rows, cols)
			then
				goto continue
			end

			local diagonalOne = matrix[row - 1][col - 1] .. matrix[row + 1][col + 1]
			local diagonalTwo = matrix[row - 1][col + 1] .. matrix[row + 1][col - 1]

			if (diagonalOne == "MS" or diagonalOne == "SM") and (diagonalTwo == "MS" or diagonalTwo == "SM") then
				occurrences = occurrences + 1
			end

			::continue::
		end
	end

	return occurrences, countA
end

local inputFileName = "input.txt"
local matrix = readMatrix(inputFileName)
local result, countA = findOccurrences(matrix)

print("The result is", result, "A count:", countA)
