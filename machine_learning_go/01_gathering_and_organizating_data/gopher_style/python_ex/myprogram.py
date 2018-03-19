import pandas as pd

'''
It is true that, very quickly, we can write some Python code to parse this CSV
and output the maximum value from the integer column without even
knowing what types are in the data:
'''

# Define column names
cols = [
    'integercolumn',
    'stringcolumn'
]

# Read in the CSV with pandas.
data = pd.read_csv('myfile.csv', names=cols)

# Print out the maximum value in the integer column.
print(data['integercolumn'].max())