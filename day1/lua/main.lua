local solution = {}
local result_one = 0
local similarity_score = 0

function solution:compareLists()
	result_one = 0
	local repeat_count = {}

	-- Create a table to store the repeated elements from list 2
	for i = 1, #self.list2 do
		if repeat_count[self.list2[i]] == nil then
			repeat_count[self.list2[i]] = 1
		else
			repeat_count[self.list2[i]] = repeat_count[self.list2[i]] + 1
		end
	end

	for i = 1, #self.list1 do
		local num = self.list1[i]
		if self.list1[i] > self.list2[i] then
			result_one = result_one + (self.list1[i] - self.list2[i])
		else
			result_one = result_one + (self.list2[i] - self.list1[i])
		end

		-- Check if the element is repeated
		if repeat_count[num] then
			similarity_score = similarity_score + (num * repeat_count[num])
		end
	end

	print(result_one)
	print("Similarity score: " .. similarity_score)
end

function solution:splitLists(input)
	self.list1 = {}
	self.list2 = {}

	local file = io.open(input, "r")
	if not file then
		error("Could not open the file: " .. input)
	end

	-- Split the input into two lists
	for line in file:lines() do
		local num1, num2 = line:match("(%d+)%s+(%d+)")

		if num1 and num2 then
			table.insert(self.list1, tonumber(num1))
			table.insert(self.list2, tonumber(num2))
		end
	end

	file:close()

	-- Oder both lists
	table.sort(self.list1)
	table.sort(self.list2)

	-- I will consider that both lists have the same size
	self:compareLists()
end

function solution:main(input)
	self:splitLists(input)
	return result_one
end

solution:main("./input.txt")
