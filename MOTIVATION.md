# Why I Made This Repo

## Dislikes
As someone who has worked some time as a backend Node.js developer, I've grown to dislike JavaScript/TypeScript, as well as some of the abstractions that popular libraries provide.

### JavaScript
I've disliked JS because it feels like a shamelessly unsafe language. Although JS has a type system, it feels very weak. It pretty much allows you to do anything regardless of probable bugs that are hard to find. Of course, part of JS's allure is precisely this lack of restraint. Experienced programmers can train themselves to avoid these bugs and take advantage of JS's freedom to do cool things that demand more lines of code in other languages. I just personally got tired of expending so much energy in avoiding bugs. Worst of all, I wasn't even good at it. I'd often find nasty or weird programming errors in my code. And, it felt upsetting that the type checker wouldn't tell me about them. They were waiting to be discovered at runtime.

### TypeScript
TS is an amazing improvement. Unlike JS, TS readily reports nasty hidden bugs to you. In addition, it has an incredibly powerful type system that feels innovative and expressive. The only reason I don't like TS is that occasionally, one has to rely on untyped JS libraries. So, you're still at risk of having subtle yet ugly type related runtime bugs every now and then. Of course, [more and more JS libraries are declaring their types](https://github.com/DefinitelyTyped/DefinitelyTyped) every day to prevent this issue. It feels backward to me that developers are writing JS library code first and then going back and adding type definitions. It feels like type definitions are more likely to be correct if the types are written before or at the same time as the code. But, this approach is probably good enough and I'm sure the problem is decreasing in severity.

### Abstraction
Abstraction is not a bad thing in and of itself. One of the most crucial aspects of design is fine-tuning the level of abstraction, so that you can worry about implementation and business logic according to your project and your preference. 

I just personally don't generally like the level of abstraction offered by big JS/TS frameworks. The main big framework I have experience with is TypeGraphQL, so I'll talk about that one here.

When I was using TypeGraphQL, I liked that I could write my type definitions in one place. Without it, I'd have to copy-paste between my schema.gql file and my typescript type definition file. That was nice and convenient.

To do this, TypeGraphQL required many layers of code between my business logic and the hardware. Most of the time those layers of code behaved as intended. And when some code within those layers failed, it was mostly because I misused the framework API. In these cases, causes and fixes were relatively easy to find and implement using documentation.

This seems fine. But, it still bothered me that there were so many layers. The more layers there are, the more code and API connections between each layer, and thus, the greater risk for 3rd party bugs. These bugs are usually much harder to find than API misuse errors. In such cases, you're basically stuck waiting for the 3rd party to fix it, or trying to find some hacky way to avoid it yourself. This pain might outweigh the benefits of the framework. In my case, the risk of that pain felt like it outweighed the benefit of only needing one type declaration file.

Performance wise, more layers equals more code, which equals more machine instructions, which equals worse performance. If [TS is already much slower compared to Go](https://youtu.be/h7UEwBaGoVo) can't even compete.

### Alternatives
Instead of switching to a thinner framework like Express.js, I wanted to implement something even closer to hardware and learn how to structure it too.

I initially started with Rust, because it's definitely a cool and lower level language with support for HTTP applications. And, that was a fun experience. But, I eventually switched to Go because it feels a bit better suited for HTTP applications and because the equivalent libraries I was planning on using were a little bit more mature.

I have the Rust project sitting around somewhere. So I'll probably rewrite it as an even lower level HTTP server using [hyper](https://github.com/hyperium/hyper).

For now though, I hope you can learn a thing or two from the code here.
