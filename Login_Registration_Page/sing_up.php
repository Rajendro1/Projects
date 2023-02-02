<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Travel Form</title>
    <link rel="stylesheet" type="text/css" href="cssfile/sing_up_style.css">
    <link href="https://fonts.googleapis.com/css?family=Roboto|Sriracha&display=swap" rel="stylesheet">
</head>
<body>
    <div class="container">
        <h1>Welcome To The Cyber World</h3>
        <p>Enter your details and submit this form to confirm your Details </p>
        <form action="after_sing_up.php" method="POST">
            <input type="text" name="name" id="name" placeholder="Enter your name" required>
            <input type="text" name="age" id="age" placeholder="Enter your Age">
            <input type="text" name="gender" id="gender" placeholder="Enter your gender">
            <input type="email" name="email" id="email" placeholder="Enter your email" required>
            <input type="phone" name="phone" id="phone" placeholder="Enter your phone">
            <input type="password" name="password" id="password" placeholder="Enter your The Password" required>
            <!-- <input type="password" name="cpassword" id="cpassword" placeholder="Confirm your Password"> -->
            <input type="file" name="image" id="image">
            <textarea name="desc" id="desc" cols="30" rows="10" placeholder="Enter any other information here Why You Visit Our Page"></textarea>
            <button class="btn">Submit</button> 
        </form>
    </div>
</body>
</html>