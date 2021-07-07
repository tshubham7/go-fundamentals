## operation over json content
So i had this problem where I got the json response from a thrid party api.
About the json response data:
    It was bank statement data, combination of unpredicted nested dictionaries and arrays.
    I had to convert all the dates which are in string into the epoch time.
Solution:
    Solution is simple, consider this json data as a tree and you want to cover every single leaf on it.
    ALGO:
        * use recursion
        * loop if the data is a dict and check for the key 'date' if you find it do the conversion and modify the value
        * also if the key is not 'date' call the function again with the resp value
        * in the second point if it's list, iterate through it and call function on each element.
        * at the bottom just return the data.

It was simple until I did in python, I never tried this in golang, so here it is.
There are some exceptions which can be taken care of easily.
Sorry for the data, I can't share it on pubic.