const int buzzerPin = 9; // Connect the positive pin of the buzzer to pin
9
void setup() {
 pinMode(buzzerPin, OUTPUT);
}
void loop() {
 tone(buzzerPin, 1000); // Play a 1 kHz tone
 delay(100); // Duration of the tick
 noTone(buzzerPin); // Stop the tone
 delay(900); // Wait before the next tick
}