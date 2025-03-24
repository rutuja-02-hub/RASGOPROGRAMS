/ Define the LED pins
const int led1 = 13; // Built-in LED (usually on pin 13)
const int led2 = 2; // External LED connected to pin 2
void setup() {
 // Initialize both pins as output
 pinMode(led1, OUTPUT);
 pinMode(led2, OUTPUT);
}
void loop() {
 // Turn on LED1 and turn off LED2
 digitalWrite(led1, HIGH);
 digitalWrite(led2, LOW);
 delay(1000); // Wait for 1 second
 // Turn off LED1 and turn on LED2
 digitalWrite(led1, LOW);
 digitalWrite(led2, HIGH);
 delay(1000); // Wait for 1 second
}