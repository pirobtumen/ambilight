#include <Adafruit_NeoPixel.h>

#define PIN 7
#define NUMPIXELS 30

Adafruit_NeoPixel pixels = Adafruit_NeoPixel(NUMPIXELS, PIN, NEO_GRB + NEO_KHZ800);

void setup() {
  pixels.begin();
  pixels.setBrightness(50);
  Serial.begin(9600);
  pinMode(LED_BUILTIN, OUTPUT);
  digitalWrite(LED_BUILTIN, LOW);

  for(int i=0;i<NUMPIXELS;i++){
    // pixels.Color takes RGB values, from 0,0,0 up to 255,255,255
    pixels.setPixelColor(i, pixels.Color(255,255,255)); // Moderately bright green color.
    delay(35);
    pixels.show(); // This sends the updated pixel color to the hardware.
  }
}

void loop() {
  digitalWrite(LED_BUILTIN, LOW);
  String color;
  bool finishRead = false;
  while(!finishRead){
      while (!Serial.available()) {}
      
      char c = Serial.read();
      if (c != '\n') {
        color += c;
      }

      if (c == ';') {
        finishRead = true;
      }
      delay(1);
  }

  if (color.length() > 0) {
    String data[3];
    int index = 0;
    int r, g, b;
    
    for (int i = 0; i < color.length(); i++) {
      if (color[i] == ',') {
        index += 1;
      } else {
        data[index] += color[i];
      }
    }

    r = data[0].toInt();
    g = data[1].toInt();
    b = data[2].toInt();

    //Serial.println(r);
    //Serial.println(g);
    //Serial.println(b);

    for(int i=0;i<NUMPIXELS;i++){
      pixels.setPixelColor(i, pixels.Color(r,g,b));
    }
    pixels.show();
    
    digitalWrite(LED_BUILTIN, HIGH);
  }
  
  delay(10);
}
