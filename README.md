# prototyping

## description

A simple program that was created to accompany a dev.to article that is in the works

## usage

```console
git clone https://github.com/oduortoni/prototyping.git
cd prototyping
```

There are two golang projects. The first one is within the folder named oldskool and the other is within the folder called newskool. The oldskool folder is the one that contains only the scaled down version of our statistical application. The newskool contains the jokes implementation included alongside the older version.

You can run the projects by going into their respective folders.

### oldskool

While in the prototyping folder

```console
cd oldskool
go run .
```

Go to your web browser and type

```plaintext
http://localhost:3500/
```

In the form input try any text and you will get an erro page. If you want to see the pdf response, type in the text, **proto**, and wait for approximately fifteen seconds for the response. 

You will realize that while you are waiting, all you have is a blank page without any interactivity. Now try the newskool to get a better user experience.

### newskool

While in the prototyping folder

```console
cd newskool
go run .
```

Go to your web browser and type

```plaintext
http://localhost:4500/
```

In the form input try any text and you will get an error page. If you want to see the pdf response, type in the text, **proto**, and wait for approximately fifteen seconds for the response. 

The beauty of the newskool page is that while you are waiting, you can click on the joke button to have a laugh session as you wait.