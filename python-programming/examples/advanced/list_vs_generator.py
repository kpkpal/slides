a_list = [i for i in range(3)]
print(a_list)

a_generator = (i for i in range(3))
print(a_generator)
for x in a_generator:
   print(x)

