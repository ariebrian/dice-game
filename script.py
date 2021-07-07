import sys
import random

def dice_roll(dice):
    roll_res = []
    for i in range(0, dice):
        roll_res.append(random.randint(1,6)) 
    return roll_res

def dice_roll_result(players_dice):
    result = {}
    for i in range(1, len(players_dice) + 1):
        print("check roll players_dice")
        print(players_dice[i])
    
        if players_dice[i] == 0:
            result[i] = [0]
        result[i] = dice_roll(players_dice[i])
    return result

def check_player(players_dice):
    playing = 0
    for i in range(1, len(players_dice)+1):
        if players_dice[i] > 0:
            playing += 1
    if playing > 1:
        return True 
    else:
        return False

def check_winner(point):
    max = 0
    key = []
    print("check point")        
    print(point)
    for i in point:
        if point.get(i) == max:
            key.append(i)
            max = point.get(i)
        elif point.get(i) > max:
            key = [i]
            max = point.get(i)
    # max_key = max(point, key=point.get)
    if len(key) > 1:
        message =  "It's a tie between " + ' '.join(map(str, point))
    else:
        message = "Winner is " + str(key[0])
    return message

n = sys.argv

player = int(n[1])
dice = int(n[2])

print("pemain " + str(player))
print("dadu " + str(dice))

# player = 2
# dice = 2
point = {}
players_dice = {}

for d in range(1, player+1):
    players_dice[d] = dice
    point[d] = 0

print(players_dice)
print(point)

a = 1
while True:
    cont_play = check_player(players_dice)
    if cont_play == False:
        break
    print("turn " + str(a))
    roll_result = dice_roll_result(players_dice)
    print(roll_result)
    a += 1
    for i in range(1, len(roll_result) + 1):
        
        print("player " + str(i))
        print(roll_result[i])
        if roll_result[i] != []:
            for res in roll_result[i]:
                if res == 6:
                    point[i] += 1
                elif res == 1:
                    players_dice[i] -= 1
                    if i < len(players_dice):
                        players_dice[i+1] += 1
                    else:
                        players_dice[1] += 1
            else:
                continue

    print("point")
    print(point)
    print("new players_dice")
    print(players_dice)
    print("========================END TURN=====================")

winner = check_winner(point)

print("========================WINNER=====================")
print(winner)
print("==================================================")