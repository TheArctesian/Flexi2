import speech_recognition as sr
import os
from termcolor import cprint

file = open("out.txt", "w")
r = sr.Recognizer()



with sr.AudioFile('int.wav') as source:
    audio = r.record(source)


    try: 

        text=r.recognize_google(audio)
        cprint("Doing things", 'purple')
        print(text)
        file.write(text)
        file.close()
        cprint("wrote to file", 'green')
    except sr.UnknownValueError:
        cprint("Google Speech Recognition could not understand audio", 'red')
    except sr.RequestError as e:
        cprint("Could not request results from Google Speech Recognition service; {0}".format(e), 'red')
    except:
        cprint("Fucky wucky", 'red')
