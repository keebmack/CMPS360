from flask import Flask, render_template
from datetime import datetime

app = Flask(__name__)

@app.route('/')
def home():
    # Current date and time
    now = datetime.now()
    # Target date and time
    target_date = datetime(2024, 6, 21)
    # Calculate the difference between target date and now
    countdown = target_date - now
    # Pass the countdown to the template
    return render_template('index.html', countdown=countdown)

if __name__ == '__main__':
    app.run(debug=True)
