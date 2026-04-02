void ___syscall_malloc() {
	puts("Nope.");
	exit(1);
}

int main() {
	char buffer[0x40];
	int x;

	printf("Please enter key: ");
	x = scanf("%23s", buffer);
	if (x == 1) {
		___syscall_malloc();
	}
	if (buffer[1] != '2') {
		___syscall_malloc();
	}
	if (buffer[0] != '4') {
		___syscall_malloc();
	}
	fflush(/* Variable de _DYNAMIC section */);
	memset(buffer + 31, 0, 9);
	buffer[31] = '*';
	char var_8 = 0; // rbp-0x41
	*(long *)(buffer + 40) = 2; // rbp-0x18
	*(int *)(buffer + 40) = 1; // rbp-0xc
				   

	if (strlen(buffer + 31) == 8) {
	}
	long var_64 = buffer[40];
	
	if (strlen(buffer) >= buffer[40]) {
	}



}
