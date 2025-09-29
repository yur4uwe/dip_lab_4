import os
from gtts import gTTS 
import time 
import pygame

pygame.mixer.init()

def listen_command(): 
    return input("Tell your team:") 
 
def do_this_command(message): 
    message = message.lower() 
    if "hello" in message: 
        say_message("hey buddy") 
    elif "how are you" in message: 
        say_message("i am fine")
    elif "so long" in message: 
        say_message("while a friend") 
        exit() 
    else: 
        say_message("command not recognized") 


def say_message(message: str): 
    print(message) 
    voice = gTTS(message, lang="en") 
    file_voice_name = "audio/" + message.replace(" ", "_") + "_audio.mp3" 
    if not os.path.exists(file_voice_name):
        voice.save(file_voice_name)
    
    pygame.mixer.music.load(file_voice_name)
    pygame.mixer.music.play()
    while pygame.mixer.music.get_busy():
        time.sleep(0.1)
    print("Voice assistant:" + message) 

if __name__ == '__main__': 
   # print_hi('PyCharm') 
    while True: 
        command = listen_command() 
        do_this_command(command)