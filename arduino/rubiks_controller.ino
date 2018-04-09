/* Rubiks Cube code hackathon */
#include <Servo.h>


#define CRADLE_SERVO_PIN 5
#define PUSHARM_SERVO_PIN 9
#define CRADLE_SERVO_CENTER 90
#define PUSHARM_SERVO_CENTER 0

// Servo's to control the rotation of the bottom of cube
Servo cradle;
int cradlePos = CRADLE_SERVO_CENTER;

// Servot to push the cube over
Servo pushArm;


char lastAction = 'l';
unsigned long lastActionMillis = 0L;
void setup() {
  Serial.begin(9600);

  // Attach and Center Cradle Servo
  cradle.attach(CRADLE_SERVO_PIN);
  cradle.write(CRADLE_SERVO_CENTER);

  // Attack and Set push arm to star position
  pushArm.attach(PUSHARM_SERVO_PIN);
  pushArm.write(PUSHARM_SERVO_CENTER);

}

void handle(char input) {
  if(input == 'a' && cradlePos != 0){
    cradle.write(0);
    cradlePos = 0;
  } else if (input == 'd' && cradlePos != 180) {
    // Rotate cube cw
    cradle.write(180);
    cradlePos = 180;
  } else if (input == 'w') {
    pushArm.write(180);
    delay(1000);
    pushArm.write(0);
  }

  if (input == 'a' || input == 'd' || input =='w') {
    lastAction = input;
  }
}

void loop() {
  // put your main code here, to run repeatedly:
  if(Serial.available() > 0) {
    char boobs = Serial.read();
    handle(boobs);
  }
}
