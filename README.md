# CAT2
<html lang="en">

<head>
    <meta charset="utf-8" />
    <meta content="width=device-width, initial-scale=1.0" name="viewport" />
    <title>
        Graham Bradshaw Musiega - Portfolio
    </title>
    <script src="https://cdn.tailwindcss.com">
    </script>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css" rel="stylesheet" />
    <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&amp;display=swap" rel="stylesheet" />
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@700&display=swap" rel="stylesheet" />
    <style>
        body {
            font-family: 'Roboto', sans-serif;
        }
    </style>
</head>

<body class="bg-gray-100 text-gray-900">
    <!-- Header & Navigation Bar -->
    <header class="bg-white shadow-md">
        <div class="container mx-auto px-6 py-3 flex justify-between items-center">
            <a class="text-2xl font-bold text-gray-800" href="#">
                Graham Bradshaw
            </a>
            <nav class="flex space-x-4">
                <a class="text-gray-800 hover:text-gray-600" href="#about">
                    About Me
                </a>
                <a class="text-gray-800 hover:text-gray-600" href="#skills">
                    Skills
                </a>
                <a class="text-gray-800 hover:text-gray-600" href="#experience">
                    Experience
                </a>
                <a class="text-gray-800 hover:text-gray-600" href="#contact">
                    Get in Touch
                </a>
                <a class="text-gray-800 hover:text-gray-600" href="https://github.com/2020-SHAW">
                    <i class="fab fa-github">
                    </i>
                </a>
            </nav>
        </div>
    </header>
    <!-- Welcome Section -->
    <section class="bg-gray-900 text-white py-20 h-screen">
        <div class="container mx-auto px-6 text-center">
            <h1 class="text-4xl font-bold mb-4">Welcome to My Portfolio</h1>
            <p class="text-xl mb-6"><br><br>
                I am Graham Bradshaw Musiega<br> <br> a highly motivated 4th-year Computer Science student at <br><br>Mama Ngina
                University College.
            </p>
            <div class="text-lg">
                <span id="typing-effect" class="typing-effect blink"></span>
            </div>
        </div>
    </section>

    <script>
        const texts = ["SEO | Web Development | Blockchain | Leadership"];
        const typingElement = document.getElementById("typing-effect");
        let index = 0;
        let charIndex = 0;
        let currentText = '';
        let isDeleting = false;

        function type() {
            if (isDeleting) {
                currentText = texts[index].substring(0, charIndex - 1);
                charIndex--;
                typingElement.textContent = currentText;

                if (charIndex === 0) {
                    isDeleting = false; // Switch to typing mode
                    index = (index + 1) % texts.length; // Move to the next text
                    setTimeout(type, 500); // Pause before starting to type again
                    return;
                }
            } else {
                currentText = texts[index].substring(0, charIndex + 1);
                charIndex++;
                typingElement.textContent = currentText;

                if (charIndex === texts[index].length) {
                    isDeleting = true; // Switch to deleting mode
                    setTimeout(type, 1000); // Pause before starting to delete
                    return;
                }
            }
            
            setTimeout(type, isDeleting ? 100 : 150); // Adjust typing and deleting speed
        }

        type(); // Start the typing effect
    </script>

    <style>
        body {
            font-family: 'Montserrat', sans-serif; /* Use Montserrat font */
        }

        .typing-effect {
            border-right: 2px solid; /* Cursor effect */
            white-space: nowrap;
            overflow: hidden;
            display: inline-block;
        }

        .blink {
            animation: blink-caret 0.75s step-end infinite;
        }

        @keyframes blink-caret {
            from, to {
                border-color: transparent;
            }
            50% {
                border-color: white; /* Change to match your text color */
            }
        }
    </style>
    <!-- About Me Section -->
    <section class="py-20 h-screen" id="about">
        <div class="container mx-auto px-6">
            <div class="flex flex-col md:flex-row items-center">
                <div class="md:w-1/3">
                    <img alt="Professional profile image of Graham Bradshaw Musiega" class="rounded-full shadow-md"
                        height="300"
                        src="https://photos.app.goo.gl/UFUxn8iMVK73go8j6"
                        width="300" />
                </div>
                <div class="md:w-2/3 md:pl-10 mt-6 md:mt-0">
                    <h2 class="text-3xl font-bold mb-4">
                        About Me
                    </h2>
                    <p class="text-lg mb-4">
                        I am a 4th-year Computer Science student at Mama Ngina University College with a focus on
                        technology, leadership, and entrepreneurship. I have a strong technical background and a passion
                        for innovation. My long-term goal is to drive change in the tech industry through teamwork and
                        innovative solutions.
                    </p>
                    <p class="text-lg">
                        I am passionate about web development, blockchain technology, and digital marketing. I believe
                        in the power of technology to transform lives and am committed to using my skills to make a
                        positive impact.
                    </p>
                </div>
            </div>
        </div>
    </section>
    <!-- Skills Section -->
    <section class="bg-gray-200 py-20" id="skills">
        <div class="container mx-auto px-6">
            <h2 class="text-3xl font-bold text-center mb-10">
                Skills
            </h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                <div class="bg-white p-6 rounded-lg shadow-md text-center">
                    <i class="fas fa-search text-4xl text-gray-800 mb-4">
                    </i>
                    <h3 class="text-xl font-bold mb-2">
                        SEO Optimization
                    </h3>
                    <p>
                        Expertise in SEO strategies and digital marketing.
                    </p>
                </div>
                <div class="bg-white p-6 rounded-lg shadow-md text-center">
                    <i class="fas fa-code text-4xl text-gray-800 mb-4">
                    </i>
                    <h3 class="text-xl font-bold mb-2">
                        Web Development
                    </h3>
                    <p>
                        Proficient in HTML, CSS, JavaScript, and React.
                    </p>
                </div>
                <div class="bg-white p-6 rounded-lg shadow-md text-center">
                    <i class="fas fa-lock text-4xl text-gray-800 mb-4">
                    </i>
                    <h3 class="text-xl font-bold mb-2">
                        Blockchain
                    </h3>
                    <p>
                        Knowledge of blockchain technologies and smart contracts.
                    </p>
                </div>
                <div class="bg-white p-6 rounded-lg shadow-md text-center">
                    <i class="fas fa-users text-4xl text-gray-800 mb-4">
                    </i>
                    <h3 class="text-xl font-bold mb-2">
                        Leadership
                    </h3>
                    <p>
                        Experience in team management and leadership roles.
                    </p>
                </div>
            </div>
        </div>
    </section>
    <!-- Experience Section -->
    <section class="py-20" id="experience">
        <div class="container mx-auto px-6">
            <h2 class="text-3xl font-bold text-center mb-10">
                Experience
            </h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div class="bg-white p-6 rounded-lg shadow-md">
                    <h3 class="text-xl font-bold mb-2">
                        SEO &amp; Digital Marketing
                    </h3>
                    <p>
                        Expertise in SEO, optimization strategies, and digital marketing.
                    </p>
                </div>
                <div class="bg-white p-6 rounded-lg shadow-md">
                    <h3 class="text-xl font-bold mb-2">
                        Web Development
                    </h3>
                    <p>
                        Experience in building modern, responsive websites with technologies like HTML, CSS, JavaScript,
                        and React.
                    </p>
                </div>
                <div class="bg-white p-6 rounded-lg shadow-md">
                    <h3 class="text-xl font-bold mb-2">
                        Blockchain Development
                    </h3>
                    <p>
                        Knowledge of blockchain technologies, smart contracts, and decentralized applications.
                    </p>
                </div>
                <div class="bg-white p-6 rounded-lg shadow-md">
                    <h3 class="text-xl font-bold mb-2">
                        Leadership &amp; Team Management
                    </h3>
                    <p>
                        Leadership roles, team-building experiences, and organizational or academic leadership
                        positions.
                    </p>
                </div>
            </div>
        </div>
    </section>
    <!-- Get in Touch Section -->
    <section class="bg-gray-200 py-20" id="contact">
        <div class="container mx-auto px-6">
            <h2 class="text-3xl font-bold text-center mb-10">
                Get in Touch
            </h2>
            <form action="/contact" class="max-w-lg mx-auto bg-white p-8 rounded-lg shadow-md" method="POST">
                <div class="mb-4">
                    <label class="block text-gray-700" for="name">Name</label>
                    <input
                        class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                        id="name" name="name" type="text" />
                </div>
                <div class="mb-4">
                    <label class="block text-gray-700" for="email">Email</label>
                    <input
                        class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                        id="email" name="email" type="email" />
                </div>
                <div class="mb-4">
                    <label class="block text-gray-700" for="message">Message</label>
                    <textarea
                        class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                        id="message" name="message"></textarea>
                </div>
                <button class="w-full bg-gray-800 text-white py-2 rounded-lg hover:bg-gray-700"
                    type="submit">Submit</button>
            </form>
        </div>
    </section>
    <!-- Footer -->
    <footer class="bg-gray-900 text-white py-6">
        <div class="container mx-auto px-6 text-center">
            <p>
                © 2023 Graham Bradshaw Musiega. All rights reserved.
            </p>
            <div class="flex justify-center space-x-4 mt-4">
                <a class="text-white hover:text-gray-400" href="">
                    <i class="fab fa-linkedin">
                    </i>
                </a>
                <a class="text-white hover:text-gray-400" href="https://github.com/2020-SHAW">
                    <i class="fab fa-github">
                    </i>
                </a>
            </div>
        </div>
    </footer>
</body>

</html>
