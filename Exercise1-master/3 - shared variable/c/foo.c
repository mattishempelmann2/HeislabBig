// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t lock;

// Note the return type: void*
void *incrementingThreadFunction(void *arg)
{
    for (int h = 0; h < 1000000; h++)
    {
        pthread_mutex_lock(&lock);
        i++;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}

void *decrementingThreadFunction(void *arg)
{
    for (int k = 0; k < 1000000; k++)
    {
        pthread_mutex_lock(&lock);
        i--;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}

void *foo(void *arg)
{
    printf("%s\n", (char *)arg);
    return NULL;
}

int main()
{
    // TODO:
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?

    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`

    pthread_t thread1, thread2;
    pthread_mutex_init(&lock, NULL);
    pthread_create(&thread1, NULL, incrementingThreadFunction, (void *)"thread1 is running");
    pthread_create(&thread2, NULL, decrementingThreadFunction, (void *)"thread2 is running");

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);
    pthread_mutex_destroy(&lock);
    printf("The magic number is: %d\n", i);
    return 0;
}
