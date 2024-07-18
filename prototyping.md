# Prototyping Software Given Time Constraints

## Introduction

Once, I attended a Google Developer Groups event packed with activities. One standout was the hackathon where six teams aimed to create unique solutions on a given theme. Our team faced a challenge: our app needed to use a specific API, but we had little time left to build a prototype. We were strangers with no common programming language, we had two options: miss out on the top prize by resorting to a basic PowerPoint presentation or try to complete a whole day's worth of work in the space of an hour.

Then, a breakthrough: despite only skimming the API documentation, we realized we knew enough to prove our app could work. Just like in math class, where understanding a component's shape lets you make educated guesses about it, we knew the API's basic function. So, we hardcoded a few examples to demonstrate how our app could use the API if fully developed.

Recently, I got roped into another hackathon-like event at work. It got me thinking: what if I face a similar situation where I need to use an API that's down? Do I wait or assume it works and build my idea around that assumption? This article shares my experience and advice for anyone in the same boat.

## Our present situation

Our company has an app that crunches user-provided data and emails out a PDF report. Now, we're proposing a cool new addition: a random jokes feature to entertain users while they wait for their results. We've got just two hours to build a prototype for a quick presentation.

Here's the plan: We'll focus on creating a simple prototype that shows off how the jokes can seamlessly fit into our app alongside the data analysis. Our main goal is to get a basic version of both functionalities up and running. With such a short timeframe, we need to manage our time wisely. Most of our effort will go into setting up the basic structure and making sure the jokes feature functions as intended. We'll keep things straightforward for now—using a set of prewritten jokes and maybe faking the data analysis part.

Our prototype will highlight how users interact with the jokes and start their data analysis. We'll create slides to back up our presentation, but the star of the show will be the working prototype itself. It's all about demonstrating the value of this new feature clearly and effectively.

During this rapid prototyping phase, we'll treat any technical uncertainties as mysteries to be solved later. The key is to show that everything works together smoothly rather than ironing out every last detail right now.

By keeping things simple and focusing on functionality, we'll nail this presentation on time and make sure our new feature shines. This approach lets us deliver a compelling demo while respecting the limits of time and tech know-how.

## implementation

### avoid network requests at all costs

If you open your terminal and type:

```console
curl https://official-joke-api.appspot.com/random_joke | jq
```

you’ll get a random joke in JSON format. JSON is how the remote site represents its jokes data.

```json
{
  "type": "general",
  "setup": "Why is peter pan always flying?",
  "punchline": "Because he neverlands",
  "id": 54
}
```

 while this endpoint could give us jokes, we might find a better option later or decide to manage jokes differently. Since time is crucial, instead of building a complex solution to fetch jokes every time, we could create a local data format that matches the JSON structure of the jokes. This way, we can prove our concept works with the data format we expect.  Later on, connecting to a remote server will be straightforward. This approach saves us from rushing into a solution that might cause more problems than it solves. All we need at this point is to understand and prepare the shape of our data. When it comes to our application, the fact that we are not actually making network requests may be abstracted away from the web page that needs that data. As a result, the front end can create the user interface based on the shape of the data we expect and still remain relevant if and when we decide to actually implement the part where we make the network requests.

 Failing to make the network request may look like a lazy solution but when faced with the constraint of time, we can not begin digging through the API's documentation in order to utilize it. All we have to do is to figure out what kind of data we expect and then provide a mock of that data within our server. During the presentation stage, we can then limit the user's input so that they can only work with a subset of the whole. A subset which we have actually implemented.

### create a simplified version of the existing features

If you can remember, our main application deals with statistical data. When we are prototyping a feature, we not only want to include the feature, we need to show that the feature will fit in seamlessly without having to sacrifice all existing features. It for this reason alone that we need to provide a scaled down version of our statistical feature. How do we go about it? Instead of creating an actual program that can do the complex statistics, we could just create a single example that can be returned based on a predetermined input. Since the statistical feature is already up and running, we could simply feed it a random input then get the results back. The result is in the form of a PDF document. We can therefore create a server that, for example, has a form to accept the same input, then each time the user enters the same input we used to derive that output(PDF file), we simply serve it as a response. Any other input will simply get a response indicating that this is but an example.

Should I implement all the major features? The answer to this will be a cryptic no. While we are trying to introduce a feature within a working ecosystem, we do not expect it to interact with every other component of the larger application. It is therefore best to avoid over-engineering a prototype. Only include what might be affected by the new changes. If you can include that without tampering with how the major feature interacts with other features, then there is no reason why you should simulate every other feature.

Maybe an example will help. Imagine having an operational vehicle with a lighting system, a transmission system and every other system. If you were to decide that your suspension was not flashy enough without the new shade of green you discovered, then, unless all the other parts were painted using a red color that will not allow the green color, then you do not need to extend your painting activities to the braking system and the body of the car. If the green color fits perfectly on your suspensions and your suspension fits perfectly with your body, then it is perfectly logical to assume that the only thing you need to match the new shade is the suspension and not the whole vehicle. I know it is not a perfect example but it highlights the fact that you do not need to do an overhaul in order to introduce a simple change.

## Example

The example will be written in the Go language.

The repository is at