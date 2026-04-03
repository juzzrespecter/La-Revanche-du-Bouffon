#include <stdio.h>
#include <string.h>

int main() {
	char buffer[20]; // 0x80, carga en 0x7a
	char input[108]; //0x6c

	buffer[120] = 0;
	buffer[6] = '_';
	buffer[7] = '_';
	buffer[8] = 's';
	buffer[9] = 't';
	buffer[10] = 'a';
	buffer[11] = 'c';
	buffer[12] = 'k';
	buffer[13] = '_';
	buffer[14] = 'c';
	buffer[15] = 'h';
	buffer[16] = 'e';
	buffer[17] = 'c';
	buffer[18] = 'k';
	buffer[19] = 0;

	printf("Please enter key: ");
	scanf("%s", input);

	if (strcmp(input, buffer + 6) == 0) {
		printf("Good job.\n");
	} else {
		printf("Nope.\n");
	}
	return 0;
}
