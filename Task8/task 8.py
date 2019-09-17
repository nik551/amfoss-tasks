from PIL import Image


from pytesseract import image_to_string
img=Image.open('C:/Users/dell/Desktop/2+22.png')
text=image_to_string(img)
print(text)
print(eval(text))
