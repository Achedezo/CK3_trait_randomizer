import random

#Variables, can be customized
min_num_of_traits = 3 
max_num_of_traits = 6
chance_of_genetic = 80 # in percent
max_num_of_genetic = 1
min_age = 16
max_age = 30
min_stat = 9 # For cost calculations, we calculate base stats between 0 and 20
max_stat = 12
education_level_weights=[15,30,30,10,5]
genetics_level_weights=[1,2,4,4,2,1]
sex_orientation_weights=[15,3,3,3]

#Game data
stats_cost = [0,2,4,6,8,12,16,20,24,31,38,45,52,63,74,85,96,113,130,147,164]
age_cost = [0,2,4,6,8,10,12,14,16,18,22,24,27,29,31,33,40,42,48,51,58,60,66,66,67,67,67,67,67,66,66,65,64,62,61,59,57,55,53,50,48,45,42,38,35,31,27,23,19,14,10,10,10,10,10,11,11,11,11,11,6,6,6,6,6,6,6,6,6,6,0]
education_cost = [0,20,40,80,150]

educations = [["Naive Appeaser", "Adequate Bargainer", "Charismatic Negotiator", "Grey Eminence","Virtuoso Arbitrator"],
             ["Misguided Warrior", "Tough Soldier", "Skilled Tactician", "Brilliant Strategist","Exalted Warlord"],
             ["Indulgent Wastrel", "Thrifty Clerk", "Fortune Builder", "Midas Touched","Golden Sovereign"],
             ["Amateurish Plotter", "Flamboyant Trickster", "Intricate Webweaver", "Elusive Shadow","Conniving Puppetmaster"],
             ["Conscientious Scribe", "Insightful Thinker", "Astute Intellectual", "Mastermind Philosopher","Erudite Oracle"]]

traits = [[["Chaste",20], ["Lustful",25]],
          [["Temperate",40], ["Gluttonous",20]],
          [["Generous",20], ["Greedy",30]],
          [["Diligent",40], ["Lazy",-10]],
          [["Patient",30], ["Impatient",25]],
          [["Humble",20], ["Arrogant",20]],
          [["Ambitious",40], ["Content",20]],
          [["Just",40], ["Arbitrary",30]],
          [["Brave",40], ["Craven",-10]],
          [["Calm",25], ["Wrathful",30]],
          [["Forgiving",25], ["Vengeful",30]],
          [["Honest",20], ["Deceitful",30]],
          [["Zealous",30], ["Cynical",30]],
          [["Gregarious",30], ["Shy",-10]],
          [["Trusting",10], ["Paranoid",-10]],
          [["Compassionate",10], ["Callous",40], ["Sadistic",40]],
          [["Fickle",25], ["Stubborn",30], ["Eccentric",15]]
          ]

genetic = [[["Hideous",-30],["Ugly",-20],["Homely",-10],["Comely",40],["Pretty / Handsome",80],["Beautiful",120]],
           [["Imbecile",-45],["Stupid",-30],["Slow",-15],["Quick",80],["Intelligent",160],["Genius",240]],
           [["Feeble",-45],["Frail",-30],["Delicate",-15],["Hale",60],["Robust",120],["Amazonian / Herculean",180]]
           ]

stats = ["Diplomacy", "Martial", "Stewardship", "Intrigue", "Learning","Prowess"]

# Character Creation
cost = []

    # Age
sex = random.choice(["Male","Female"])
sex_orient = random.choices(["Heterosexual","Homosexual","Bisexual","Asexual"],sex_orientation_weights)
age = random.randint(min_age, max_age)
if age > 70:
    cost.append(0)
else:
    cost.append(age_cost[age])

    # Education
education_level = random.choices([0,1,2,3,4],education_level_weights)
education_trait = random.randint(0,4)
education_final = educations[education_trait][education_level[0]]
cost.append(education_cost[education_level[0]])

    #traits
num_of_traits = random.randint(min_num_of_traits, max_num_of_traits)
traits_char = []
for _ in range(num_of_traits):
    duo = random.choice(traits) #can be a trio now, doesn't change anything
    traits.remove(duo)
    chosen = random.choice(duo)
    traits_char.append(chosen[0])
    cost.append(chosen[1])
    
    #genetic
genetics=[]
for _ in range(max_num_of_genetic):
    if random.randint(1, 100) <= chance_of_genetic:
        category = random.choice(genetic)
        genetic.remove(category)
        chosen = random.choices(category,genetics_level_weights)
        genetics.append(chosen[0][0])
        cost.append(chosen[0][1])

#stats
stats_character = []
for stat in stats:
    stat_value = random.randint(min_stat, max_stat)
    stats_character.append(stat_value)
    cost.append(stats_cost[stat_value])
print()

# Printing and Formating
print("Age: " + str(age))
print("Sex: " + sex)
print("Sexual Orientation: "+ sex_orient[0])
print("Education: "+ education_final)
print("Personality Traits:")
for i in traits_char:
    print("\t"+i)
if genetics:
    print("Genetics:")
    for i in genetics:
        print("\t"+i)
print("Statistics:")
print("\tDiplomacy:   "+ str(stats_character[0]) +
      "\n\tMartial:     "+ str(stats_character[1]) +
      "\n\tStewardship: "+ str(stats_character[2]) +
      "\n\tIntrigue:    "+ str(stats_character[3]) +
      "\n\tLearning:    "+ str(stats_character[4]) +
      "\n\tProwess:     "+ str(stats_character[5])
      )
print("Cost: "+str(sum(cost)))
if sum(cost)<=400:
    print("Achievements Enabled")
else:
    print("Achievements Disabled")