import random

#can be cusomized
min_num_of_traits = 3 
max_num_of_traits = 6
chance_of_genetic = 10 # in percent
max_num_of_genetic = 1
min_age = 16
max_age = 30
min_stat = 2
max_stat = 6

educations = [["Naive Appeaser", "Underhanded Rogue", "Charismatic Negotiator", "Grey Eminence"],
             ["Misguided Warrior", "Tough Soldier", "Skilled Tactician", "Brilliant Strategist"],
             ["Indulgent Wastrel", "Thrifty Clerk", "Fortune Builder", "Midas Touched"],
             ["Amateurish Plotter", "Flamboyant Schemer", "Intricate Webweaver", "Elusive Shadow"],
            ["Detached Priest", "Dutiful Cleric", "Scholarly Theologian", "Mastermind Theologian"]]

traits = [["Chaste", "Lustful"],
          ["Temperate", "Gluttonous"],
          ["Charitable", "Greedy"],
          ["Diligent", "Slothful"],
          ["Patient", "Wroth"],
          ["Humble", "Proud"],
          ["Ambitious", "Content"],
          ["Just", "Arbitrary"],
          ["Brave", "Craven"],
          ["Kind", "Cruel"],
          ["Honest", "Deceitful"],
          ["Zealous", "Cynical"],
          ["Gregarious", "Shy"],
          ["Trusting", "Paranoid"]]

genetic = [["Attractive", "Ugly"],
           ["Strong", "Weak"],
           ["Genius", "Quick", "Slow", "Imbecile"]]

stats = ["Diplomacy", "Martial", "Stewardship", "Intrigue", "Learning"]


num_of_traits = random.randint(min_num_of_traits, max_num_of_traits)

#age
print("Age: " + str(random.randint(min_age, max_age)) + "\n")

#stats
for stat in stats:
    print(stat + ": " + str(random.randint(min_stat, max_stat)))
print()

#education
print(random.choice(random.choice(educations))) 

#traits
for _ in range(num_of_traits):
    duo = random.choice(traits)
    traits.remove(duo)
    print(random.choice(duo))
    
#genetic
for _ in range(max_num_of_genetic):
    if random.randint(1, 100) <= chance_of_genetic:
        duo = random.choice(genetic)
        genetic.remove(duo)
        print(random.choice(duo))
    