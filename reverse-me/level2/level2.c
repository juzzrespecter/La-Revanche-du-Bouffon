#include <stdio.h>

void no() {
	puts("Nope.");
	exit(1);
}

void ok() {
}

int main() {
	int bool_41;
	char atoi_buffer[3]; //0x39 - 0x37
	int int0_36; //0x36
	char buffer[23]; // 0x35 - 0x1E
	char buffer_2[9]; // 0x1d
	

	printf("Please enter key: ");
	int value = scanf("%23s", buffer);
	if (value != 1) {
		no();
	}
	if (buffer[1] != 0x30) {
		no();
	}
	if (buffer[0] != 0x30) {
		no();
	}
	fflush();
	memset(buffer_2, 0, 9);
	buffer_2[0] = 0x64;
	int int0_36 = 0; //0x36 
	int int0_14 = 2; //0x14 
	int int0_10 = 1; //0x10 
.main_221:
	while (true) {
		int x = strlen(buffer_2);
		if (x < 0x8)
			// jmp to main+286
		// ebp - 0x48 = ebp - 0x14
		int y = strlen(buffer);
		if (y >= /*ebp - 0x48 */) {
			break ; // jmp main+378
		}
.main_302:
	//  ebp - 0x40
	//  ebp - 0x14	
		atoi_buffer[0] = buffer[c];
		atoi_buffer[1] = buffer[c + 1];
		atoi_buffer[2] = buffer[c + 2];

		buffer_2[d] = atoi(atoi_buffer);
		int0_14 = int0_14 + 3;
		int0_10 = int0_10 + 1;
	} // jmp main+221
.main_378:
	buffer[int0_10] = 0;
	if (strcmp(buffer, "delabere") == 0) {
		ok();
	}
	no();
	return 0;
}
