package ingredients

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

func init() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel("trace")
	} else {
		log.SetLevel("info")
	}
}

func BenchmarkParse(b *testing.B) {
	log.SetLevel("error")
	fileToGet := ts[0].URL
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewFromFile(fileToGet)
	}
}

func ExampleReadme() {
	r, _ := NewFromURL("https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/")
	fmt.Println(r.IngredientList())
	// Output:
	// 1 cup butter (salted)
	// 1 cup sugar (white)
	// 1 cup brown sugar (light)
	// 2 tsp vanilla (pure)
	// 2 whole eggs
	// 3 cups flour (all purpose)
	// 1 tsp baking soda
	// 1/2 tsp baking powder
	// 1 tsp salt (sea)
	// 2 cups chocolate chips
}

type URLIngredients struct {
	URL         string
	Ingredients []string
}

var ts = []URLIngredients{
	{
		"https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies",
		[]string{"1 1/4 tsp salt", "3/4 tsp baking soda", "3/4 cup butter", "1 cup brown sugar", "1/4 cup sugar", "1 whole egg", "2 whole egg yolks", "2 tsp vanilla", "6 oz chocolate chips"},
	},
	{
		"https://pinchofyum.com/the-best-soft-chocolate-chip-cookies",
		[]string{"8 tablespoons butter", "1/2 cup sugar", "1/4 cup brown sugar", "1 teaspoon vanilla", "1 whole egg", "1 1/2 cups flour", "1/2 teaspoon baking soda", "1/4 teaspoon salt", "3/4 cup chocolate chips"},
	},
	{
		"https://www.modernhoney.com/the-best-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup brown sugar", "1/2 cup sugar", "2 whole eggs", "2 teaspoons vanilla", "2 3/4 cups flour", "1 teaspoon cornstarch", "3/4 teaspoon baking soda", "3/4 teaspoon salt", "2 1/2 cups chocolate chips"},
	},
	{
		"https://laurenslatest.com/actually-perfect-chocolate-chip-cookies/",
		[]string{"3/4 cup butter", "1 cup brown sugar", "1/2 cup sugar", "1 whole egg yolk", "2 teaspoons vanilla", "2 cups flour", "1/2 teaspoon baking soda", "1/2 teaspoon salt", "1 cup chocolate chips", "1 cup chocolate chips"},
	},
	{
		"https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup sugar", "1 cup brown sugar", "2 whole eggs", "2 teaspoons vanilla", "1 teaspoon baking soda", "2 teaspoons water", "1/2 teaspoon salt", "3 cups flour", "2 cups chocolate chips", "1 cup walnuts"},
	},
	{
		"https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup sugar", "1 cup brown sugar", "2 tsp vanilla", "2 whole eggs", "3 cups flour", "1 tsp baking soda", "1/2 tsp baking powder", "1 tsp salt", "2 cups chocolate chips"},
	},
	{
		"https://cakebycourtney.com/soft-chewy-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup sugar", "1 cup brown sugar", "2 whole eggs", "1 teaspoon vanilla", "3 1/2 cups flour", "1 teaspoon baking soda", "1 teaspoon baking powder", "1 teaspoon salt", "2 cups chocolate chips"},
	},
	{
		"https://www.foodnetwork.com/recipes/dave-lieberman/noodle-kugel-recipe-1946564",
		[]string{"1/2 pound egg noodles", "1/2 whole butter", "1 pound cottage cheese", "2 cups sour cream", "1/2 cup sugar", "6 whole eggs", "1 teaspoon cinnamon", "1/2 cup raisins"},
	},
	{
		"https://cooking.nytimes.com/recipes/12320-apple-pie",
		[]string{"2 tablespoons butter", "2 1/2 pounds apples", "1/4 teaspoon allspice", "1/2 teaspoon cinnamon", "1/4 teaspoon salt", "3/4 cup sugar", "2 tablespoons flour", "2 teaspoons cornstarch", "1 tablespoon apple cider vinegar", "1 whole pie dough", "1 whole egg"},
	},
	{
		"www.cooksillustrated.com/recipes/11519-indian-butter-chicken-murgh-makhani",
		[]string{"2 tablespoons butter", "1/2 whole onion", "3 whole garlic", "2 teaspoons ginger", "2 teaspoons serrano chile", "1 1/2 teaspoons garam masala", "1/2 teaspoon coriander", "1/4 teaspoon cumin", "1/4 teaspoon pepper", "3/4 cup water", "1/4 cup tomato paste", "1 1/2 teaspoons sugar", "1 teaspoon salt", "1/2 cup heavy cream", "1 pound chicken thighs", "1/4 cup greek yogurt", "1 1/2 tablespoons cilantro", "11 g sugar"},
	},
}

func TestTable(t *testing.T) {
	for i, t0 := range ts {
		log.Info(i, t0.URL)
		fileToGet := t0.URL
		fileToGet = strings.TrimPrefix(fileToGet, "https://")
		if string(fileToGet[len(fileToGet)-1]) == "/" {
			fileToGet += "index.html"
		}
		fileToGet = path.Join("testing", "sites", fileToGet)
		r, err := NewFromFile(fileToGet)
		assert.Nil(t, err)
		ingredients := make([]string, len(r.IngredientList().Ingredients))
		for i, ing := range r.IngredientList().Ingredients {
			ingredients[i] = fmt.Sprintf("%s %s %s", AmountToString(ing.Measure.Amount), ing.Measure.Name, ing.Name)
		}
		assert.Equal(t, t0.Ingredients, ingredients)
	}
}

func ExampleExtractJSON() {

	var htmlString string
	htmlString = `
<html>
<head>
</head>
<body>
<script>

	[{"@context":"http://schema.org","@type":"Recipe","mainEntityOfPage":true,"name":"Chocolate Chip Cookies","url":"https://www.foodnetwork.com/recipes/food-network-kitchen/chocolate-chip-cookies-recipe4-2011856","headline":"Chocolate Chip Cookies","description":"This is such an easy chocolate chip cookie. No special equipment, no creaming -- a perfect cookie to do with kids. We love how versatile this dough is, too. It makes an awesome rocky road bar cookie.","author":[{"@type":"Person","name":"Food Network Kitchen","url":"https://www.foodnetwork.com/profiles/talent/food-network-kitchen"}],"image":{"@type":"ImageObject","url":"https://food.fnr.sndimg.com/content/dam/images/food/fullset/2006/12/11/0/12Days_Chocolate_Chip_Cooki.jpg.rend.hgtvcom.406.305.suffix/1391394585748.jpeg","height":"406","width":"305"},"datePublished":"2016-11-16T00:45:02.056-05:00","dateModified":"2015-01-22T12:56:46.193-05:00","publisher":{"@type":"Organization","name":"Food Network","url":"https://www.foodnetwork.com","logo":{"@type":"ImageObject","url":"https://food.fnr.sndimg.com/etc/clientlibs/assets/images/food/fn-logo-flat-60x60.png","height":"60","width":"60"}},"keywords":"Easy Dessert Recipes,Dessert,Easy,Easy Baking,Baking,American,Cookie,Chocolate Chip Cookie,Christmas,Holiday,Chocolate,Kid-Friendly","cookTime":"P0Y0M0DT0H15M0.000S","prepTime":"P0Y0M0DT0H20M0.000S","totalTime":"P0Y0M0DT1H5M0.000S","recipeIngredient":["1/2 cup (1 stick) unsalted butter","3/4 cup packed dark brown sugar","3/4 cup sugar","2 large eggs","1 teaspoon pure vanilla extract","1 (12-ounce) bag semisweet chocolate chips, or chunks","2 1/4 cups all-purpose flour","3/4 teaspoon baking soda","1 teaspoon fine salt"],"recipeInstructions":[{"@type":"HowToStep","text":"Evenly position 2 racks in the middle of the oven and preheat to 375 degrees F. (on convection setting if you have it.) Line 2 baking sheets with parchment paper or silicone sheets. (If you only have 1 baking sheet, let it cool completely between batches.)"},{"@type":"HowToStep","text":"Put the butter in a microwave safe bowl, cover and microwave on medium power until melted. (Alternatively melt in a small saucepan.) Cool slightly. Whisk the sugars, eggs, butter and vanilla in a large bowl until smooth."},{"@type":"HowToStep","text":"Whisk the flour, baking soda and salt in another bowl. Stir the dry ingredients into the wet ingredients with a wooden spoon; take care not to over mix. Stir in the chocolate chips or chunks."},{"@type":"HowToStep","text":"Scoop heaping tablespoons of the dough onto the prepared pans. Wet hands slightly and roll the dough into balls. Space the cookies about 2-inches apart on the pans. Bake, until golden, but still soft in the center, 12 to 16 minutes, depending on how chewy or crunchy you like your cookies. Transfer hot cookies with a spatula to a rack to cool. Serve."},{"@type":"HowToStep","text":"Store cookies in a tightly sealed container for up to 5 days."},{"@type":"HowToStep","text":"For a Rocky Road Bar:"},{"@type":"HowToStep","text":"Lightly butter a 9 by 13-inch baking pan. Make the batter as per cookie recipe and fold in 1 cup chopped walnuts along with the chocolate chips. Spread batter in prepared pan. Bake until the edges are light brown and the batter sets, about 45 minutes. Cool slightly and cover surface with 4 cups marshmallows and 1 cup chocolate chips. Broil at least 8 inches from the heat until marshmallows turn golden brown, about 2 minutes. (Keep an eye on the marshmallows, and turn the pan frequently--they go from golden to char in a wink.) Cool, cut and serve."}],"aggregateRating":{"@type":"AggregateRating","ratingValue":4.2,"reviewCount":536},"recipeYield":"30 cookies","review":[{"@type":"Review","author":{"@type":"Person","name":"JACQUELINE M."},"reviewRating":{"@type":"Rating","ratingValue":5,"worstRating":"1","bestRating":"5"},"reviewBody":"If you cream the butter and sugar first until light and fluffy, then the egg and vanilla, then flour, they come out exactly as shown.  ","datePublished":"2018-07-28"},{"@type":"Review","author":{"@type":"Person","name":"Neysa S."},"reviewRating":{"@type":"Rating","ratingValue":1,"worstRating":"1","bestRating":"5"},"reviewBody":"I wish I would have read the reviews on here before baking these cookies. Normally I am good about doing that. Anyway, I followed the recipe to a T! My first batch came out super hard. I was really disappointed. So, what I did with the rest was put it in a 9x13 glass baking dish. I coated it with butter then flour. Reduced the heat to 350 and baked until top was golden brown. I now have soft and chewy chocolate chip cookie bars. Hope this helps for anyone out there. ","datePublished":"2016-12-14"},{"@type":"Review","author":{"@type":"Person","name":"Shayna T."},"reviewRating":{"@type":"Rating","ratingValue":1,"worstRating":"1","bestRating":"5"},"reviewBody":" Too many chocolate chips and was not a cakey cookie I was going for. I will not use recipe again.","datePublished":"2016-07-09"},{"@type":"Review","author":{"@type":"Person","name":"Anonymous"},"reviewRating":{"@type":"Rating","ratingValue":1,"worstRating":"1","bestRating":"5"},"reviewBody":"Very bland, too much flour. These were so bad that I ate one and threw the rest away. Even when under baked, they were still very crunchy -- almost too hard. I do not recommend this recipe.","datePublished":"2015-11-03"},{"@type":"Review","author":{"@type":"Person","name":"lothman10"},"reviewBody":"I made these and they taste like bread! Bad cookies but very good bread","datePublished":"2018-01-17"},{"@type":"Review","author":{"@type":"Person","name":"Anonymous"},"reviewRating":{"@type":"Rating","ratingValue":1,"worstRating":"1","bestRating":"5"},"reviewBody":"Not good. I followed the recipe exactly and my cookies came out like chocolate chip hockey pucks. They were too dry, lacked any flavor, and was a general waste of ingredients.","datePublished":"2017-12-10"},{"@type":"Review","author":{"@type":"Person","name":"Anonymous"},"reviewRating":{"@type":"Rating","ratingValue":4,"worstRating":"1","bestRating":"5"},"reviewBody":"I saw the reviews oon here and decided to make a few changes when I made mine. 1st thing was using a stand up mixer. Next was instead of melting the butter, i softened it and creamed it wirh the sugars and vanilla before adding the eggs. Last was adding a about an equal ammount of baking powder to compliment the baking soda (I don't measure these and the salt exactly when making cookies). They turned out really well like this, though they didn't spread too much they still made wonderful cookies. ","datePublished":"2017-08-02"},{"@type":"Review","author":{"@type":"Person","name":"Anonymous"},"reviewRating":{"@type":"Rating","ratingValue":5,"worstRating":"1","bestRating":"5"},"reviewBody":"Recipe works great just ignore the over mixing and mix very well. Then mash or knead the crumbly dough till it looks like the cookie dough you like. Butter parchment paper place on baking pan. Don't form balls with dough it won't spread. Instead shape cookie to desired size like a burger Patti and bake for 10-11 minutes max for soft cookie.","datePublished":"2017-04-20"},{"@type":"Review","author":{"@type":"Person","name":"sbrnkm"},"reviewBody":"I don't know what these positive reviews are on about, these cookies are TERRIBLE. There must be some mistake with the recipe, either it's supposed to be 2 sticks of butter instead of 1 or it's calling for way too much flour.<br /><br />Don't ignore the negative comments on this one as I did lol, I used my last cup of chocolate chips to make these as a surprise for my mom when she gets home and I'm stuck with what is essentially a crunchy flavorless cake with chocolate chips tossed in :/<br /><br />They were the exact same shape after baking as they were when I placed them on the pan, they did not cook evenly at all, and some of the chocolate was burnt despite the rest looking underdone with some crunchy overbaked parts on top. Really really bad stuff. <br />","datePublished":"2017-01-20"},{"@type":"Review","author":{"@type":"Person","name":"charmaine d."},"reviewBody":"loved this recipe.. one word of advice I would not melt the butter. just let it sit on the counter over night and use the soft butter.","datePublished":"2016-10-26"}],"recipeCuisine":"american","recipeCategory":"dessert"}, {"@context":"http://schema.org","@type":"BreadcrumbList","itemListElement":[{"@type":"ListItem","position":1,"item":{"@id":"https://www.foodnetwork.com","name":"Home"}},{"@type":"ListItem","position":2,"item":{"@id":"https://www.foodnetwork.com/recipes","name":"Recipes"}}]}, {"@context":"http://schema.org","@type":"ItemList","itemListElement":[{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/recipes","position":1,"name":"Recipes"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/shows","position":2,"name":"Shows"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/profiles","position":3,"name":"Chefs"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/restaurants","position":4,"name":"Restaurants"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/videos","position":5,"name":"Videos"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/magazine","position":6,"name":"Magazine"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/features/articles/sweepstakes-and-contests","position":7,"name":"Sweeps"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/shows/tv-schedule","position":8,"name":"TV Schedule"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/fn-dish","position":9,"name":"Blog"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/site/apps","position":10,"name":"Apps"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/shows/a-z","position":11,"name":"Shows A-Z"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/profiles/talent","position":12,"name":"Chefs A-Z"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/healthy","position":13,"name":"Healthy"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/holidays-and-parties","position":14,"name":"Party Food"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/grilling","position":15,"name":"Grilling"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/how-to/packages/shopping","position":16,"name":"Shop"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/shows/chopped","position":17,"name":"Chopped"},{"@type":"SiteNavigationElement","url":"https://www.foodnetwork.com/shows/the-great-food-truck-race","position":18,"name":"The Great Food Truck Race"}]}]
</script>
</body>
</html>
`

	r, err := NewFromString(htmlString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.IngredientList())
	// Output:
	// 	1/2 cup butter (unsalted)
	// 3/4 cup brown sugar (packed dark)
	// 3/4 cup sugar
	// 2 whole eggs
	// 1 teaspoon vanilla (pure)
	// 1 whole chocolate chips
	// 2 1/4 cups flour (all purpose)
	// 3/4 teaspoon baking soda
	// 1 teaspoon salt (fine)

}

func ExampleExtractJSON2() {

	var htmlString string
	htmlString = `
<html>
<head>
</head>
<body>
<script>
{
	    "@context": "https://schema.org",
	    "@graph": [
	    {
	        "@type": "WebSite",
	        "@id": "https://thesaltymarshmallow.com/#website",
	        "url": "https://thesaltymarshmallow.com/",
	        "name": "The Salty Marshmallow",
	        "potentialAction":
	        {
	            "@type": "SearchAction",
	            "target": "https://thesaltymarshmallow.com/?s={search_term_string}",
	            "query-input": "required name=search_term_string"
	        }
	    },
	    {
	        "@type": "ImageObject",
	        "@id": "https://thesaltymarshmallow.com/best-banana-bread-recipe/#primaryimage",
	        "url": "https://thesaltymarshmallow.com/wp-content/uploads/2018/11/banana-bread6.jpg",
	        "width": 680,
	        "height": 1020,
	        "caption": "banana bread sliced"
	    },
	    {
	        "@type": "WebPage",
	        "@id": "https://thesaltymarshmallow.com/best-banana-bread-recipe/#webpage",
	        "url": "https://thesaltymarshmallow.com/best-banana-bread-recipe/",
	        "inLanguage": "en-US",
	        "name": "Best Banana Bread Recipe - The Salty Marshmallow",
	        "isPartOf":
	        {
	            "@id": "https://thesaltymarshmallow.com/#website"
	        },
	        "primaryImageOfPage":
	        {
	            "@id": "https://thesaltymarshmallow.com/best-banana-bread-recipe/#primaryimage"
	        },
	        "datePublished": "2018-11-30T10:00:48+00:00",
	        "dateModified": "2019-01-16T22:19:25+00:00",
	        "author":
	        {
	            "@id": "https://thesaltymarshmallow.com/#/schema/person/5213878e4bb22cced598c7323c139c2f"
	        },
	        "description": "Best Banana Bread Recipe is so easy to make and super soft and moist!\u00a0 The very best way to use up overripe bananas this bread is tender and packed full of flavor!"
	    },
	    {
	        "@type": ["Person"],
	        "@id": "https://thesaltymarshmallow.com/#/schema/person/5213878e4bb22cced598c7323c139c2f",
	        "name": "Nichole",
	        "image":
	        {
	            "@type": "ImageObject",
	            "@id": "https://thesaltymarshmallow.com/#authorlogo",
	            "url": "https://secure.gravatar.com/avatar/691f101a10157925374a2f59f8a9027a?s=96&d=mm&r=g",
	            "caption": "Nichole"
	        },
	        "sameAs": []
	    },
	    {
	        "@context": "http://schema.org/",
	        "@type": "Recipe",
	        "name": "Best Banana Bread Recipe",
	        "author":
	        {
	            "@type": "Person",
	            "name": "Nichole"
	        },
	        "description": "Best Banana Bread Recipe is so easy to make and super soft and moist!&nbsp; The very best way to use up overripe bananas this bread is tender and packed full of flavor!",
	        "datePublished": "2018-11-30T04:00:48+00:00",
	        "image": ["https://thesaltymarshmallow.com/wp-content/uploads/2018/11/banana-bread6.jpg", "https://thesaltymarshmallow.com/wp-content/uploads/2018/11/banana-bread6-500x500.jpg", "https://thesaltymarshmallow.com/wp-content/uploads/2018/11/banana-bread6-500x375.jpg", "https://thesaltymarshmallow.com/wp-content/uploads/2018/11/banana-bread6-480x270.jpg"],
	        "recipeYield": "1 Loaf",
	        "prepTime": "PT15M",
	        "cookTime": "PT50M",
	        "totalTime": "PT65M",
	        "recipeIngredient": ["1 Stick Butter", "3 Large Ripe Bananas", "2 Large Eggs", "1 teaspoon Vanilla Extract", "2 Cups All Purpose Flour", "1 Cup Granulated Sugar", "1 teaspoon Baking Soda", "1/2 teaspoon salt", "1/2 teaspoon cinnamon"],
	        "recipeInstructions": [
	        {
	            "@type": "HowToStep",
	            "text": "Preheat oven to 350 degrees.  Spray a loaf pan with non-stick cooking spray or grease with butter and set aside."
	        },
	        {
	            "@type": "HowToStep",
	            "text": "Add the stick of butter to a large bowl and microwave for 1 minute, or until melted."
	        },
	        {
	            "@type": "HowToStep",
	            "text": "Add the bananas to the same bowl and mash with a fork."
	        },
	        {
	            "@type": "HowToStep",
	            "text": "Add the vanilla extract and egg to the bowl and use the same fork to mash and stir until no yellow streaks of egg remain."
	        },
	        {
	            "@type": "HowToStep",
	            "text": "In a second large bowl whisk together the flour, sugar, baking soda, salt, and cinnamon."
	        },
	        {
	            "@type": "HowToStep",
	            "text": "Add the dry ingredients to the wet ingredients and mix together with a spatula just until combined."
	        },
	        {
	            "@type": "HowToStep",
	            "text": "Pour the batter into prepared loaf pan and bake for 45-55 minutes until a toothpick inserted in the center of the bread comes out clean."
	        }],
	        "recipeCategory": ["Dessert"],
	        "recipeCuisine": ["American"],
	        "keywords": "banana bread, banana bread recipe, best banana bread",
	        "nutrition":
	        {
	            "@type": "NutritionInformation",
	            "servingSize": "1 g",
	            "calories": "252 kcal",
	            "carbohydrateContent": "40 g",
	            "proteinContent": "4 g",
	            "fatContent": "9 g",
	            "saturatedFatContent": "5 g",
	            "cholesterolContent": "55 mg",
	            "sodiumContent": "270 mg",
	            "fiberContent": "2 g",
	            "sugarContent": "21 g"
	        },
	        "aggregateRating":
	        {
	            "@type": "AggregateRating",
	            "ratingValue": "4.97",
	            "ratingCount": "88"
	        },
	        "@id": "https://thesaltymarshmallow.com/best-banana-bread-recipe/#recipe",
	        "isPartOf":
	        {
	            "@id": "https://thesaltymarshmallow.com/best-banana-bread-recipe/#article"
	        },
	        "mainEntityOfPage": "https://thesaltymarshmallow.com/best-banana-bread-recipe/#webpage"
	    }]
	}</script>
</body>
</html>
`
	r, err := NewFromString(htmlString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.IngredientList())
	// Output:
	// 1 whole butter
	// 3 whole bananas
	// 2 whole eggs
	// 1 teaspoon vanilla
	// 2 cups flour (all purpose)
	// 1 cup sugar (granulated)
	// 1 teaspoon baking soda
	// 1/2 teaspoon salt
	// 1/2 teaspoon cinnamon

}
