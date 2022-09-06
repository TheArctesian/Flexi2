newRooms = []
newTeachers = []
temp = 0
t = ''
for x in rm: 
    if x == temp:
       continue 
    elif x !=temp: 
        newRooms.append(x)
        temp = x

for x in teachers: 
    if x == t:
       continue 
    else: 
        newTeachers.append(x)
        t = x

print(newRooms)
print(newTeachers)
print("Array length =", len(newRooms))
print("Array length =", len(newTeachers))
