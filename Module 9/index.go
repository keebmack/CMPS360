package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <title>Home</title>
		    <style>
		        body {
		            font-family: Arial, sans-serif;
		            margin: 0;
		            padding: 0;
		            background-color: #f5f5f5;
		        }
		        .container {
		            max-width: 800px;
		            margin: 20px auto;
		            padding: 20px;
		            background-color: #fff;
		            border-radius: 5px;
		            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		        }
		        h1 {
		            color: #333;
		        }
		        p {
		            line-height: 1.6;
		        }
		        .navigation {
		            margin-top: 20px;
		        }
		        .navigation a {
		            text-decoration: none;
		            color: #333;
		            margin-right: 20px;
		        }
		        .profile-image {
		            max-width: 100%;
		            height: auto;
		            border-radius: 5px;
		            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		            max-width: 150px; /* Adjust the maximum width as needed */
		        }
		        .hobby-image {
		            max-width: 100%;
		            height: auto;
		            border-radius: 5px;
		            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		            max-width: 200px; /* Adjust the maximum width as needed */
		        }
		    </style>
		</head>
		<body>
		
		<div class="container">
		    <h1>Welcome to My Website!</h1>
		    <p>
		        The website of Ryan Kebe, Powered by GO
		    </p>
		    <div class="navigation">
		        <a href="/about">About Me</a>
		        <a href="/hobby">Hobbies and Interests</a>
		    </div>
		    <div>
		    </div>
		</div>
		
		</body>
		</html>`)
	})

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <title>About Me</title>
		    <style>
		        body {
		            font-family: Arial, sans-serif;
		            margin: 0;
		            padding: 0;
		            background-color: #f5f5f5;
		        }
		        .container {
		            max-width: 800px;
		            margin: 20px auto;
		            padding: 20px;
		            background-color: #fff;
		            border-radius: 5px;
		            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		        }
		        h1 {
		            color: #333;
		        }
		        p {
		            line-height: 1.6;
		        }
		        .profile-image {
		            max-width: 100%;
		            height: auto;
		            margin-bottom: 20px;
		            border-radius: 5px;
		            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
					max-width: 200px;
		        }
		    </style>
		</head>
		<body>
		
		<div class="container">
		    <h1>About Me</h1>
		    <p>
		        My name is Ryan Kebe and I am 30 years old and currently live in Meadville with my fiance, Kate.
		    </p>
		    <p>
		        I am a Broadband Technician for Armstrong. I have a pug named Rico, and he is my best friend. I am originally from Greensburg and I will soon be relocating to Zelienople. I look forward to moving back into the Pittsburgh Area!
		    </p>
		    <img src="/images/rico.jpg" alt="Rico" class="profile-image">
		    <img src="/images/niagara.jpg" alt="Niagara Falls" class="profile-image">
			<img src="https://chambermaster.blob.core.windows.net/images/customers/712/members/190/logos/MEMBER_PAGE_HEADER/Armstrong_Logo_no_tag.png" alt="Armstrong Logo" class="profile-image">
		</div>
		
		</body>
		</html>`)
	})

	http.HandleFunc("/hobby", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <title>Hobbies and Interests</title>
		    <style>
		        body {
		            font-family: Arial, sans-serif;
		            margin: 0;
		            padding: 0;
		            background-color: #f5f5f5;
		        }
		        .container {
		            max-width: 800px;
		            margin: 20px auto;
		            padding: 20px;
		            background-color: #fff;
		            border-radius: 5px;
		            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		        }
		        h1 {
		            color: #333;
		        }
		        p {
		            line-height: 1.6;
		        }
		        .hobby {
		            margin-bottom: 20px;
		        }
		        .hobby-image {
		            max-width: 100%;
		            height: auto;
		            border-radius: 5px;
		            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		            max-width: 200px;
		        }
		    </style>
		</head>
		<body>

		<div class="container">
		    <h1>Hobbies and Interests</h1>
		    <div class="hobby">
		        <h2>Pittsburgh Steelers</h2>
		        <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/de/Pittsburgh_Steelers_logo.svg/768px-Pittsburgh_Steelers_logo.svg.png" alt="Pittsburgh Steelers" class="hobby-image">
		    </div>
		    <div class="hobby">
		        <h2>Pittsburgh Penguins</h2>
		        <img src="https://i.pinimg.com/originals/d1/80/e1/d180e1d5c75c42703f216be319b1bc54.jpg" alt="Pittsburgh Penguins" class="hobby-image">
		    </div>
			<div class="hobby">
		        <h2>Video Games</h2>
		        <img src="https://upload.wikimedia.org/wikipedia/en/8/8d/Dark_Souls_Cover_Art.jpg" alt="Dark Souls" class="hobby-image">
				<img src="https://upload.wikimedia.org/wikipedia/en/b/b9/Elden_Ring_Box_art.jpg" alt="Elden Ring" class="hobby-image">
				<img src="https://upload.wikimedia.org/wikipedia/en/thumb/5/53/MorrowindCOVER.jpg/220px-MorrowindCOVER.jpg" alt="The Elder Scrolls" class="hobby-image">
				<img src="https://upload.wikimedia.org/wikipedia/en/a/af/Fallout.jpg" alt="Fallout" class="hobby-image">
				<img src="https://image.api.playstation.com/vulcan/ap/rnd/202302/2321/3098481c9164bb5f33069b37e49fba1a572ea3b89971ee7b.jpg" alt="Baldur's Gate 3" class="hobby-image">
		    </div>
		</div>

		</body>
		</html>`)
	})

	fmt.Println("Server listening on Port 8080...")
	http.ListenAndServe(":8080", nil)
}
