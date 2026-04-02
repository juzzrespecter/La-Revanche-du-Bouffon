#include <stdio.h>

int main() {
	char buffer[20]; // 0x80
	char input[108]; //0x6c

	buffer[120] = 0;
	*(int *)(buffer + 6) = "__st";
	*(int *)(buffer + 10) = "ack_";
	*(int *)(buffer + 14) = "chec";
	buffer[16] = 'k';
	buffer[17] = 0;

	printf("Please enter key: ");
	scanf("%s", input);


}
