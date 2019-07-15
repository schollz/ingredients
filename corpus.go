package recipeingredients

var herbMap = map[string]struct{}{
	"ajwain":               {},
	"alligator pepper":     {},
	"allspice":             {},
	"amchoor":              {},
	"angelica":             {},
	"anise":                {},
	"aonori":               {},
	"aromatic ginger":      {},
	"asafoetida":           {},
	"basil":                {},
	"bay leaf":             {},
	"black cardamom":       {},
	"black mustard":        {},
	"black peppercorn":     {},
	"boldo":                {},
	"bolivian coriander":   {},
	"borage":               {},
	"brazilian pepper":     {},
	"brown mustard":        {},
	"bunium persicum":      {},
	"camphor":              {},
	"caraway":              {},
	"cardamom":             {},
	"cassia":               {},
	"cayenne pepper":       {},
	"celery powder":        {},
	"celery seed":          {},
	"charoli":              {},
	"chenpi":               {},
	"chervil":              {},
	"chili pepper":         {},
	"chives":               {},
	"cicely":               {},
	"cilantro":             {},
	"cinnamon":             {},
	"clove":                {},
	"coriander leaf":       {},
	"coriander seed":       {},
	"cress":                {},
	"cubeb":                {},
	"culantro":             {},
	"cumin":                {},
	"curry leaf":           {},
	"dill":                 {},
	"dill seed":            {},
	"dried lime":           {},
	"east asian pepper":    {},
	"epazote":              {},
	"fennel":               {},
	"fenugreek":            {},
	"fingerroot":           {},
	"garlic":               {},
	"garlic chives":        {},
	"ginger":               {},
	"golpar":               {},
	"grains of paradise":   {},
	"grains of selim":      {},
	"greater galangal":     {},
	"green peppercorn":     {},
	"hemp":                 {},
	"hoja santa":           {},
	"holy basil":           {},
	"horseradish":          {},
	"houttuynia cordata":   {},
	"hyssop":               {},
	"indian bay leaf":      {},
	"jimbu":                {},
	"juniper berry":        {},
	"kinh gioi":            {},
	"kokum":                {},
	"korarima":             {},
	"lavender":             {},
	"lemon balm":           {},
	"lemon grass":          {},
	"lemon myrtle":         {},
	"lemon verbena":        {},
	"lesser galangal":      {},
	"limnophila aromatica": {},
	"liquorice":            {},
	"litsea cubeba":        {},
	"long pepper":          {},
	"lovage":               {},
	"mace":                 {},
	"mahlab":               {},
	"mango-ginger":         {},
	"marjoram":             {},
	"mastic":               {},
	"mint":                 {},
	"mitsuba":              {},
	"mugwort":              {},
	"nigella":              {},
	"nigella sativa":       {},
	"njangsa":              {},
	"nutmeg":               {},
	"oregano":              {},
	"paprika":              {},
	"parsley":              {},
	"perilla":              {},
	"peruvian pepper":      {},
	"pomegranate seed":     {},
	"poppy seed":           {},
	"radhuni":              {},
	"rose":                 {},
	"rosemary":             {},
	"rue":                  {},
	"saffron":              {},
	"sage":                 {},
	"salt":                 {},
	"sansho":               {},
	"sarsaparilla":         {},
	"sassafras":            {},
	"savory":               {},
	"sesame":               {},
	"shiso":                {},
	"sichuan pepper":       {},
	"sorrel":               {},
	"star anise":           {},
	"sumac":                {},
	"tamarind":             {},
	"tarragon":             {},
	"tasmanian pepper":     {},
	"thai basil":           {},
	"thyme":                {},
	"tonka bean":           {},
	"turmeric":             {},
	"uzazi":                {},
	"vanilla":              {},
	"vietnamese coriander": {},
	"voatsiperifery":       {},
	"wasabi":               {},
	"white mustard":        {},
	"white peppercorn":     {},
	"woodruff":             {},
	"yuzu":                 {},
	"zedoary":              {},
	"zereshk":              {},
	"zest":                 {},
}

var fruitMap = map[string]struct{}{
	"apple":             {},
	"apricot":           {},
	"avocado":           {},
	"banana":            {},
	"bell pepper":       {},
	"bilberry":          {},
	"blackberry":        {},
	"blackcurrant":      {},
	"blood orange":      {},
	"blueberry":         {},
	"boysenberry":       {},
	"breadfruit":        {},
	"canary melon":      {},
	"cantaloupe":        {},
	"cherimoya":         {},
	"cherry":            {},
	"chili pepper":      {},
	"clementine":        {},
	"cloudberry":        {},
	"coconut":           {},
	"cranberry":         {},
	"cucumber":          {},
	"currant":           {},
	"damson":            {},
	"date":              {},
	"dragonfruit":       {},
	"durian":            {},
	"eggplant":          {},
	"elderberry":        {},
	"feijoa":            {},
	"fig":               {},
	"goji berry":        {},
	"gooseberry":        {},
	"grape":             {},
	"grapefruit":        {},
	"guava":             {},
	"honeydew":          {},
	"huckleberry":       {},
	"jackfruit":         {},
	"jambul":            {},
	"jujube":            {},
	"kiwi fruit":        {},
	"kumquat":           {},
	"lemon":             {},
	"lime":              {},
	"loquat":            {},
	"lychee":            {},
	"mandarine":         {},
	"mango":             {},
	"mulberry":          {},
	"nectarine":         {},
	"nut":               {},
	"olive":             {},
	"orange":            {},
	"pamelo":            {},
	"papaya":            {},
	"passionfruit":      {},
	"peach":             {},
	"pear":              {},
	"persimmon":         {},
	"physalis":          {},
	"pineapple":         {},
	"plum":              {},
	"pomegranate":       {},
	"pomelo":            {},
	"purple mangosteen": {},
	"quince":            {},
	"raisin":            {},
	"rambutan":          {},
	"raspberry":         {},
	"redcurrant":        {},
	"rock melon":        {},
	"salal berry":       {},
	"satsuma":           {},
	"star fruit":        {},
	"strawberry":        {},
	"tamarillo":         {},
	"tangerine":         {},
	"ugli fruit":        {},
	"watermelon":        {},
}

var vegetableMap = map[string]struct{}{
	"acorn squash":        {},
	"alfalfa sprout":      {},
	"amaranth":            {},
	"anise":               {},
	"artichoke":           {},
	"arugula":             {},
	"asparagus":           {},
	"aubergine":           {},
	"azuki bean":          {},
	"banana squash":       {},
	"basil":               {},
	"bean sprout":         {},
	"beet":                {},
	"black bean":          {},
	"black-eyed pea":      {},
	"bok choy":            {},
	"borlotti bean":       {},
	"broad beans":         {},
	"broccoflower":        {},
	"broccoli":            {},
	"brussels sprout":     {},
	"butternut squash":    {},
	"cabbage":             {},
	"calabrese":           {},
	"caraway":             {},
	"carrot":              {},
	"cauliflower":         {},
	"cayenne pepper":      {},
	"celeriac":            {},
	"celery":              {},
	"chamomile":           {},
	"chard":               {},
	"chayote":             {},
	"chickpea":            {},
	"chives":              {},
	"cilantro":            {},
	"collard green":       {},
	"corn":                {},
	"corn salad":          {},
	"courgette":           {},
	"cucumber":            {},
	"daikon":              {},
	"delicata":            {},
	"dill":                {},
	"eggplant":            {},
	"endive":              {},
	"fennel":              {},
	"fiddlehead":          {},
	"frisee":              {},
	"garlic":              {},
	"gem squash":          {},
	"ginger":              {},
	"green bean":          {},
	"green pepper":        {},
	"habanero":            {},
	"herbs and spice":     {},
	"horseradish":         {},
	"hubbard squash":      {},
	"jalapeno":            {},
	"jerusalem artichoke": {},
	"jicama":              {},
	"kale":                {},
	"kidney bean":         {},
	"kohlrabi":            {},
	"lavender":            {},
	"leek ":               {},
	"legume":              {},
	"lemon grass":         {},
	"lentils":             {},
	"lettuce":             {},
	"lima bean":           {},
	"mamey":               {},
	"mangetout":           {},
	"marjoram":            {},
	"mung bean":           {},
	"mushroom":            {},
	"mustard green":       {},
	"navy bean":           {},
	"new zealand spinach": {},
	"nopale":              {},
	"okra":                {},
	"onion":               {},
	"oregano":             {},
	"paprika":             {},
	"parsley":             {},
	"parsnip":             {},
	"patty pan":           {},
	"pea":                 {},
	"pinto bean":          {},
	"potato":              {},
	"pumpkin":             {},
	"radicchio":           {},
	"radish":              {},
	"rhubarb":             {},
	"rosemary":            {},
	"runner bean":         {},
	"rutabaga":            {},
	"sage":                {},
	"scallion":            {},
	"shallot":             {},
	"skirret":             {},
	"snap pea":            {},
	"soy bean":            {},
	"spaghetti squash":    {},
	"spinach":             {},
	"squash ":             {},
	"sweet potato":        {},
	"tabasco pepper":      {},
	"taro":                {},
	"tat soi":             {},
	"thyme":               {},
	"topinambur":          {},
	"tubers":              {},
	"turnip":              {},
	"wasabi":              {},
	"water chestnut":      {},
	"watercress":          {},
	"white radish":        {},
	"yam":                 {},
	"zucchini":            {},
}

var corpusIngredients = []string{" can cream of chicken soups ",
	" nonhydrogenated margarines ",
	" can cream of chicken soup ",
	" nonhydrogenated margarine ",
	" pomegranate concentrates ",
	" cream of mushroom soups ",
	" pomegranate concentrate ",
	" cocktail pumpernickels ",
	" cream of mushroom soup ",
	" graham cracker crumbss ",
	" miniature marshmallows ",
	" raspberry vinaigrettes ",
	" salt and black peppers ",
	" balsamic vinaigrettes ",
	" blueberry blackberrys ",
	" butterscotch puddings ",
	" chanterelle mushrooms ",
	" cocktail pumpernickel ",
	" decaffeinated coffees ",
	" graham cracker crumbs ",
	" heavy whipping creams ",
	" jagermeister liqueurs ",
	" meringue buttercreams ",
	" miniature marshmallow ",
	" monosodium glutamates ",
	" raspberry blackberrys ",
	" raspberry vinaigrette ",
	" salt and black pepper ",
	" tamarind concentrates ",
	" vegetable shortenings ",
	" vinaigrette dressings ",
	" worcestershire sauces ",
	" apple cider vinegars ",
	" assorted decorations ",
	" balsamic vinaigrette ",
	" benedictine liqueurs ",
	" blueberry blackberry ",
	" butterscotch pudding ",
	" chanterelle mushroom ",
	" chipotle mayonnaises ",
	" confectioners sugars ",
	" crystallized gingers ",
	" cultured buttermilks ",
	" decaffeinated coffee ",
	" elderflower cordials ",
	" elderflower liqueurs ",
	" guanciale pancettums ",
	" heavy whipping cream ",
	" hellmann mayonnaises ",
	" jagermeister liqueur ",
	" jerusalem artichokes ",
	" limeade concentrates ",
	" limnophila aromaticas ",
	" marinated artichokes ",
	" meringue buttercream ",
	" monosodium glutamate ",
	" orange flower waters ",
	" parmigiano reggianos ",
	" pepperoncini peppers ",
	" portabella mushrooms ",
	" portobello mushrooms ",
	" raspberry blackberry ",
	" shallot vinaigrettes ",
	" sweetened cranberrys ",
	" szechuan peppercorns ",
	" tamarind concentrate ",
	" vegetable shortening ",
	" vietnamese corianders ",
	" vinaigrette dressing ",
	" worcestershire sauce ",
	" amontillado sherrys ",
	" appenzeller cheeses ",
	" apple cider vinegar ",
	" assorted decoration ",
	" assorted vegetables ",
	" balsamic reductions ",
	" benedictine liqueur ",
	" bucatini spaghettis ",
	" buffalo mozzarellas ",
	" buttermilk biscuits ",
	" butterscotch sauces ",
	" cabernet sauvignons ",
	" californium walnuts ",
	" chartreuse liqueurs ",
	" cheesecake fillings ",
	" chicken tenderloins ",
	" chilled buttermilks ",
	" chipotle mayonnaise ",
	" chocolate frostings ",
	" clementine segments ",
	" confectioners sugar ",
	" corkscrew macaronis ",
	" crystalized gingers ",
	" crystallized ginger ",
	" cultured buttermilk ",
	" desiccated coconuts ",
	" elderflower cordial ",
	" elderflower liqueur ",
	" guanciale pancettum ",
	" hellmann mayonnaise ",
	" jerusalem artichoke ",
	" jerusalem artichokes ",
	" limeade concentrate ",
	" limnophila aromatica ",
	" limoncello liqueurs ",
	" luxardo maraschinos ",
	" maraschino liqueurs ",
	" marinated artichoke ",
	" nasturtium blossoms ",
	" new zealand spinachs ",
	" orange flower water ",
	" parmigiano reggiano ",
	" peppermint schnapps ",
	" pepperoncini pepper ",
	" portabella mushroom ",
	" portobello mushroom ",
	" portuguese sausages ",
	" precooked polentums ",
	" pumpernickel breads ",
	" quality mayonnaises ",
	" romanesco broccolis ",
	" rotisserie chickens ",
	" shallot vinaigrette ",
	" shichimi togarashis ",
	" sichuan peppercorns ",
	" sourdough baguettes ",
	" spaghetti linguines ",
	" spinach fettuccines ",
	" strawberry gelatins ",
	" sweetened cranberry ",
	" szechuan peppercorn ",
	" vegetable bouillons ",
	" vegetarian sausages ",
	" venison tenderloins ",
	" vietnamese coriander ",
	" white wine vinegars ",
	" yellowtail snappers ",
	" amontillado sherry ",
	" andouille sausages ",
	" appenzeller cheese ",
	" apple concentrates ",
	" assorted mushrooms ",
	" assorted vegetable ",
	" balsamic reduction ",
	" barbecued chickens ",
	" black peppercornss ",
	" blackberry brandys ",
	" blanched hazelnuts ",
	" bolivian corianders ",
	" breakfast sausages ",
	" broccoli flowerets ",
	" bucatini spaghetti ",
	" buffalo mozzarella ",
	" buttermilk biscuit ",
	" butterscotch chips ",
	" butterscotch sauce ",
	" cabernet sauvignon ",
	" californium walnut ",
	" candy thermometers ",
	" caramelized onions ",
	" cellophane noodles ",
	" champagne vinegars ",
	" chartreuse liqueur ",
	" cheese tortellinis ",
	" cheesecake filling ",
	" chicken tenderloin ",
	" chilled buttermilk ",
	" chilled champagnes ",
	" chimichurri sauces ",
	" chinkiang vinegars ",
	" chocolate biscuits ",
	" chocolate frosting ",
	" chocolate liqueurs ",
	" chocolate puddings ",
	" chocolate shavings ",
	" chocolate soymilks ",
	" chunky applesauces ",
	" clementine segment ",
	" cointreau liqueurs ",
	" corkscrew macaroni ",
	" cornstarch sifteds ",
	" cranberry liqueurs ",
	" cream of mushrooms ",
	" crisco shortenings ",
	" crystalized ginger ",
	" cultivated mussels ",
	" desiccated coconut ",
	" digestive biscuits ",
	" dijon vinaigrettes ",
	" emmentaler cheeses ",
	" european cucumbers ",
	" fingerling potatos ",
	" flatbread crackers ",
	" framboise liqueurs ",
	" gingerbread spices ",
	" gingersnap cookies ",
	" gorgonzola cheeses ",
	" grains of paradises ",
	" granular fructoses ",
	" granulated garlics ",
	" green bell peppers ",
	" hardwood charcoals ",
	" horseradish sauces ",
	" hothouse cucumbers ",
	" houttuynia cordatas ",
	" hulled strawberrys ",
	" hungarian paprikas ",
	" imitation vanillas ",
	" italian seasonings ",
	" japanese cucumbers ",
	" japanese eggplants ",
	" jerusalem artichoke ",
	" juice concentrates ",
	" kefalotyri cheeses ",
	" kewpie mayonnaises ",
	" lebanese cucumbers ",
	" lemon vinaigrettes ",
	" limoncello liqueur ",
	" lowfat buttermilks ",
	" lowfat mayonnaises ",
	" luxardo maraschino ",
	" malaguetum peppers ",
	" manzanilla sherrys ",
	" maraschino cherrys ",
	" maraschino liqueur ",
	" marshmallow creams ",
	" marshmallow cremes ",
	" marshmallow fluffs ",
	" mascarpone cheeses ",
	" mccormick paprikas ",
	" mexican chocolates ",
	" mozzarella cheeses ",
	" nasturtium blossom ",
	" neufchatel cheeses ",
	" new zealand spinach ",
	" nonfat buttermilks ",
	" nonfat mayonnaises ",
	" parmesan pecorinos ",
	" peekytoe crabmeats ",
	" peppermint schnapp ",
	" pickled vegetables ",
	" pickling cucumbers ",
	" pomegranate juices ",
	" pomegranate syrups ",
	" porterhouse steaks ",
	" portuguese sausage ",
	" precooked polentum ",
	" pumpernickel bread ",
	" pumpkin pie spices ",
	" quality mayonnaise ",
	" raspberry liqueurs ",
	" raspberry vinegars ",
	" red pepper flakess ",
	" romanesco broccoli ",
	" rotisserie chicken ",
	" shaken buttermilks ",
	" shichimi togarashi ",
	" shiitake mushrooms ",
	" shoestring potatos ",
	" shortbread cookies ",
	" shortcrust pastrys ",
	" sichuan peppercorn ",
	" slender asparaguss ",
	" smoked mozzarellas ",
	" sourdough baguette ",
	" sourdough croutons ",
	" sourdough starters ",
	" spaghetti linguine ",
	" spinach fettuccine ",
	" strawberry gelatin ",
	" strawberry yogurts ",
	" sweetened coconuts ",
	" turkey tenderloins ",
	" unblanched almonds ",
	" vanilla flavorings ",
	" vanilla ice creams ",
	" vegetable bouillon ",
	" vegetarian sausage ",
	" venison tenderloin ",
	" vermicelli noodles ",
	" watermelon radishs ",
	" wheat orecchiettes ",
	" white horseradishs ",
	" white wine vinegar ",
	" whole wheat flours ",
	" yellow chartreuses ",
	" yellowtail snapper ",
	" additional salsas ",
	" amaretti biscuits ",
	" amaretto liqueurs ",
	" amaro montenegros ",
	" andouille sausage ",
	" apple concentrate ",
	" artichoke bottoms ",
	" artichoke heartss ",
	" assorted fillings ",
	" assorted mushroom ",
	" assorted toppings ",
	" baking chocolates ",
	" balsamic vinegars ",
	" barbecued chicken ",
	" beefsteak tomatos ",
	" berry strawberrys ",
	" black peppercorns ",
	" blackberry brandy ",
	" blackberry syrups ",
	" blanched hazelnut ",
	" bolivian coriander ",
	" bottle champagnes ",
	" breakfast cereals ",
	" breakfast radishs ",
	" breakfast sausage ",
	" broccoli floretss ",
	" broccoli floweret ",
	" buckwheat noodles ",
	" butter margarines ",
	" buttermilk breads ",
	" butternut squashs ",
	" butterscotch chip ",
	" camembert cheeses ",
	" canadian whiskeys ",
	" candy thermometer ",
	" caramelized onion ",
	" carbonated waters ",
	" cellophane noodle ",
	" champagne vinegar ",
	" cheese tortellini ",
	" cherrystone clams ",
	" chicken bouillons ",
	" chicken drumettes ",
	" chihuahua cheeses ",
	" chilled champagne ",
	" chilled proseccos ",
	" chimichurri sauce ",
	" chinese broccolis ",
	" chinese eggplants ",
	" chinkiang vinegar ",
	" chocolate biscuit ",
	" chocolate cookies ",
	" chocolate liqueur ",
	" chocolate morsels ",
	" chocolate mousses ",
	" chocolate pudding ",
	" chocolate shaving ",
	" chocolate soymilk ",
	" chocolate spreads ",
	" chunky applesauce ",
	" clarified butters ",
	" clementine juices ",
	" cocchi americanos ",
	" cocktail sausages ",
	" coconut crystalss ",
	" coconut macaroons ",
	" cointreau liqueur ",
	" compressed yeasts ",
	" cornflake cereals ",
	" cornichon pickles ",
	" cornish game hens ",
	" cornstarch sifted ",
	" cranberry liqueur ",
	" cream of mushroom ",
	" cremini mushrooms ",
	" crimini mushrooms ",
	" crisco shortening ",
	" cubanelle peppers ",
	" cultivated mussel ",
	" dandelion flowers ",
	" decorative sugars ",
	" delicatum squashs ",
	" digestive biscuit ",
	" dijon mayonnaises ",
	" dijon vinaigrette ",
	" dumpling wrappers ",
	" east asian peppers ",
	" emmentaler cheese ",
	" empanada wrappers ",
	" english cucumbers ",
	" espelette peppers ",
	" european cucumber ",
	" fingerling potato ",
	" flatbread cracker ",
	" framboise liqueur ",
	" german chocolates ",
	" ginger marmalades ",
	" gingerbread spice ",
	" gingersnap cookie ",
	" gorgonzola cheese ",
	" gorgonzola dolces ",
	" grains of paradise ",
	" granular fructose ",
	" granulated garlic ",
	" grapefruit juices ",
	" green bell pepper ",
	" green chartreuses ",
	" green peppercorns ",
	" hardwood charcoal ",
	" hazelnut liqueurs ",
	" homogenized milks ",
	" honeycrisp apples ",
	" horseradish sauce ",
	" hot pepper sauces ",
	" hothouse cucumber ",
	" houttuynia cordata ",
	" hulled strawberry ",
	" hungarian paprika ",
	" imitation vanilla ",
	" instant espressos ",
	" instant polentums ",
	" italian dressings ",
	" italian eggplants ",
	" italian meatballs ",
	" italian parmesans ",
	" italian seasoning ",
	" jalapeno chiliess ",
	" japanese cucumber ",
	" japanese eggplant ",
	" japanese pumpkins ",
	" jarlsberg cheeses ",
	" juice concentrate ",
	" kefalotyri cheese ",
	" kentucky bourbons ",
	" kewpie mayonnaise ",
	" kielbasa sausages ",
	" lebanese cucumber ",
	" lemon vinaigrette ",
	" linguica sausages ",
	" lowfat buttermilk ",
	" lowfat margarines ",
	" lowfat mayonnaise ",
	" maitake mushrooms ",
	" malaguetum pepper ",
	" manzanilla sherry ",
	" maraschino cherry ",
	" marshmallow cream ",
	" marshmallow creme ",
	" marshmallow fluff ",
	" marshmallow peeps ",
	" mascarpone cheese ",
	" mccormick paprika ",
	" mexican chocolate ",
	" mixed peppercorns ",
	" mozzarella cheese ",
	" multigrain breads ",
	" mushroom risottos ",
	" neufchatel cheese ",
	" nondairy creamers ",
	" nonfat buttermilk ",
	" nonfat mayonnaise ",
	" orange marmalades ",
	" panko breadcrumbs ",
	" pareve margarines ",
	" parmesan croutons ",
	" parmesan pecorino ",
	" parsley cilantros ",
	" pecorino toscanos ",
	" peekytoe crabmeat ",
	" pencil asparaguss ",
	" pickled cucumbers ",
	" pickled jalapenos ",
	" pickled vegetable ",
	" pickling cucumber ",
	" pomegranate juice ",
	" pomegranate seeds ",
	" pomegranate syrup ",
	" porcini mushrooms ",
	" porterhouse steak ",
	" processed cheeses ",
	" prosciutto cottos ",
	" provolone cheeses ",
	" pumpkin pie spice ",
	" purple mangosteens ",
	" raspberry liqueur ",
	" raspberry sorbets ",
	" raspberry vinegar ",
	" red pepper flakes ",
	" red wine vinegars ",
	" reposado tequilas ",
	" roasting chickens ",
	" roquefort cheeses ",
	" saffron optionals ",
	" shaken buttermilk ",
	" shaved chocolates ",
	" shiitake mushroom ",
	" shitake mushrooms ",
	" shoestring potato ",
	" shortbread cookie ",
	" shortbread crusts ",
	" shortcrust pastry ",
	" skinned hazelnuts ",
	" slender asparagus ",
	" smoked bratwursts ",
	" smoked mozzarella ",
	" smoked whitefishs ",
	" solid shortenings ",
	" sourdough crouton ",
	" sourdough starter ",
	" southern comforts ",
	" spaghetti noodles ",
	" spaghetti squashs ",
	" spanish mackerels ",
	" spinach linguines ",
	" spinach tortillas ",
	" splash grenadines ",
	" splenda granulars ",
	" square focacciums ",
	" strawberry juices ",
	" strawberry purees ",
	" strawberry sauces ",
	" strawberry syrups ",
	" strawberry yogurt ",
	" sugar substitutes ",
	" sunflower sprouts ",
	" sweetened coconut ",
	" tarragon parsleys ",
	" tarragon vinegars ",
	" tenderloin roasts ",
	" togarashi peppers ",
	" turkey pepperonis ",
	" turkey tenderloin ",
	" unblanched almond ",
	" unbleached flours ",
	" unprocessed brans ",
	" valencium oranges ",
	" vanilla flavoring ",
	" vanilla frostings ",
	" vanilla ice cream ",
	" vegan mayonnaises ",
	" vegetarian bacons ",
	" vegetarian chilis ",
	" velveetum cheeses ",
	" vermicelli noodle ",
	" watermelon juices ",
	" watermelon radish ",
	" wheat breadcrumbs ",
	" wheat fettuccines ",
	" wheat orecchiette ",
	" white grapefruits ",
	" white horseradish ",
	" white peppercorns ",
	" whole wheat flour ",
	" wildflower honeys ",
	" winter vegetables ",
	" yellow chartreuse ",
	" zucchini blossoms ",
	" additional salsa ",
	" alligator peppers ",
	" amaretti biscuit ",
	" amaretti cookies ",
	" amaretto liqueur ",
	" amaro montenegro ",
	" american cheeses ",
	" apricot gelatins ",
	" apricot liqueurs ",
	" armagnac cognacs ",
	" arrowroot flours ",
	" artichoke bottom ",
	" artichoke hearts ",
	" artichoke hearts ",
	" assorted filling ",
	" assorted topping ",
	" atlantic salmons ",
	" bacon pancettums ",
	" baking chocolate ",
	" balsamic vinegar ",
	" beef tenderloins ",
	" beefsteak tomato ",
	" berry strawberry ",
	" black peppercorns ",
	" blackberry syrup ",
	" blanched almonds ",
	" bloody mary mixs ",
	" blueberry juices ",
	" bottle champagne ",
	" bourbon whiskeys ",
	" brazilian peppers ",
	" breakfast cereal ",
	" breakfast radish ",
	" brewed espressos ",
	" broccoli florets ",
	" broiler chickens ",
	" buckwheat flours ",
	" buckwheat groats ",
	" buckwheat honeys ",
	" buckwheat noodle ",
	" burratum cheeses ",
	" butter margarine ",
	" butter optionals ",
	" buttered noodles ",
	" buttermilk bread ",
	" butternut squash ",
	" butternut squashs ",
	" buttery crackers ",
	" button mushrooms ",
	" cacao chocolates ",
	" caesar dressings ",
	" calamatum olives ",
	" camembert cheese ",
	" canadian whiskey ",
	" cannellini beans ",
	" caramel toppings ",
	" carbonated water ",
	" cerignola olives ",
	" challah brioches ",
	" champagne grapes ",
	" champagne yeasts ",
	" charentai melons ",
	" cherry tomatoess ",
	" cherrystone clam ",
	" chervil parsleys ",
	" chicken bouillon ",
	" chicken breastss ",
	" chicken carcasss ",
	" chicken drumette ",
	" chicken gizzards ",
	" chicken sausages ",
	" chickpea rinseds ",
	" chicory lettuces ",
	" chihuahua cheese ",
	" chilled prosecco ",
	" chilled seltzers ",
	" chinese broccoli ",
	" chinese cabbages ",
	" chinese eggplant ",
	" chinese mustards ",
	" chinese sausages ",
	" chipotle peppers ",
	" chocolate cookie ",
	" chocolate crusts ",
	" chocolate morsel ",
	" chocolate mousse ",
	" chocolate sauces ",
	" chocolate spread ",
	" chocolate syrups ",
	" chocolate wafers ",
	" chorizo sausages ",
	" ciabattum breads ",
	" cilantro leavess ",
	" cipolline onions ",
	" cipollini onions ",
	" clarified butter ",
	" clementine juice ",
	" cocchi americano ",
	" cocktail peanuts ",
	" cocktail sausage ",
	" coconut crystals ",
	" coconut macaroon ",
	" coconut shavings ",
	" coconut vinegars ",
	" common mushrooms ",
	" compressed yeast ",
	" cornbread crumbs ",
	" cornflake cereal ",
	" cornflake crumbs ",
	" cornichon pickle ",
	" cornish game hen ",
	" country sausages ",
	" cranberry jellys ",
	" cranberry juices ",
	" cranberry sauces ",
	" cranberry vodkas ",
	" cream of tartars ",
	" cremini mushroom ",
	" crimini mushroom ",
	" croissant doughs ",
	" crusty baguettes ",
	" cubanelle pepper ",
	" cultured butters ",
	" dandelion flower ",
	" dandelion greens ",
	" decorative sugar ",
	" delicatum squash ",
	" desired toppings ",
	" dijon mayonnaise ",
	" dumpling wrapper ",
	" east asian pepper ",
	" elephant garlics ",
	" emmental cheeses ",
	" empanada wrapper ",
	" enchilada sauces ",
	" english cucumber ",
	" english mustards ",
	" espelette pepper ",
	" espresso coffees ",
	" evaporated milks ",
	" florida avocados ",
	" flour tortillass ",
	" focaccium breads ",
	" focaccium doughs ",
	" freestone peachs ",
	" french baguettes ",
	" gelatin desserts ",
	" german chocolate ",
	" ginger marmalade ",
	" globe artichokes ",
	" golden flaxseeds ",
	" gorgonzola dolce ",
	" graham crackerss ",
	" grapefruit juice ",
	" grapefruit peels ",
	" greater galangals ",
	" green asparaguss ",
	" green chartreuse ",
	" green corianders ",
	" green peppercorn ",
	" green peppercorns ",
	" grenadine syrups ",
	" habanero peppers ",
	" halloumi cheeses ",
	" hazelnut butters ",
	" hazelnut liqueur ",
	" hazelnut spreads ",
	" hefeweizen beers ",
	" homogenized milk ",
	" honeycrisp apple ",
	" hot pepper sauce ",
	" iceberg lettuces ",
	" iceburg lettuces ",
	" instant couscous ",
	" instant espresso ",
	" instant polentum ",
	" instant tapiocas ",
	" israeli couscous ",
	" italian dressing ",
	" italian eggplant ",
	" italian fontinas ",
	" italian meatball ",
	" italian parmesan ",
	" italian parsleys ",
	" italian sausages ",
	" jalapeno chilies ",
	" jalapeno peppers ",
	" japanese pumpkin ",
	" jarlsberg cheese ",
	" kalamatum olives ",
	" kentucky bourbon ",
	" kielbasa sausage ",
	" kitchen bouquets ",
	" lamb tenderloins ",
	" lasagna noodless ",
	" lavender flowers ",
	" lemon marmalades ",
	" lime mayonnaises ",
	" lingonberry jams ",
	" linguica sausage ",
	" littleneck clams ",
	" lowfat margarine ",
	" lumpium wrappers ",
	" macaroni noodles ",
	" macintosh apples ",
	" maitake mushroom ",
	" manchego cheeses ",
	" mandarin oranges ",
	" margaritum salts ",
	" marshmallow peep ",
	" merguez sausages ",
	" mexican chorizos ",
	" mexican oreganos ",
	" milk mozzarellas ",
	" miniature bagels ",
	" mixed peppercorn ",
	" mixed vegetables ",
	" muenster cheeses ",
	" multigrain bread ",
	" muscovado sugars ",
	" mushroom risotto ",
	" mzarella cheeses ",
	" nondairy creamer ",
	" nonpareil capers ",
	" orange marmalade ",
	" oyster mushrooms ",
	" pancettum bacons ",
	" panko breadcrumb ",
	" pareve margarine ",
	" parmesan cheeses ",
	" parmesan crouton ",
	" parmesan shaveds ",
	" parsley chervils ",
	" parsley cilantro ",
	" pasteurized eggs ",
	" pattypan squashs ",
	" peach nectarines ",
	" pecorino cheeses ",
	" pecorino romanos ",
	" pecorino toscano ",
	" pencil asparagus ",
	" peppadew peppers ",
	" pickled cucumber ",
	" pickled herrings ",
	" pickled jalapeno ",
	" pickled mustards ",
	" piloncillo cones ",
	" pimiento peppers ",
	" pineapple juices ",
	" pineapple salsas ",
	" piquillo peppers ",
	" pistachio pecans ",
	" polish kielbasas ",
	" polska kielbasas ",
	" pomegranate seed ",
	" pomegranate seeds ",
	" porcini mushroom ",
	" pork tenderloins ",
	" premium tequilas ",
	" pressure cookers ",
	" processed cheese ",
	" prosciutto cotto ",
	" provolone cheese ",
	" purple mangosteen ",
	" raclette cheeses ",
	" raspberry coulis ",
	" raspberry juices ",
	" raspberry purees ",
	" raspberry sauces ",
	" raspberry sorbet ",
	" raspberry vodkas ",
	" red bell peppers ",
	" red wine vinegar ",
	" remoulade sauces ",
	" reposado tequila ",
	" rice vermicellis ",
	" roasting chicken ",
	" romaine lettuces ",
	" roquefort cheese ",
	" rosemary branchs ",
	" ruby grapefruits ",
	" saffron optional ",
	" salt and peppers ",
	" saltine crackers ",
	" saskatoon berrys ",
	" sauvignon blancs ",
	" seltzer chilleds ",
	" shaved chocolate ",
	" shaved parmesans ",
	" shishito peppers ",
	" shitake mushroom ",
	" shortbread crust ",
	" shrimp bouillons ",
	" skinned hazelnut ",
	" smoked bluefishs ",
	" smoked bratwurst ",
	" smoked kielbasas ",
	" smoked mackerels ",
	" smoked whitefish ",
	" solid shortening ",
	" sourdough breads ",
	" southern comfort ",
	" spaetzle noodles ",
	" spaghetti noodle ",
	" spaghetti squash ",
	" spaghetti squashs ",
	" spanish chorizos ",
	" spanish mackerel ",
	" spanish paprikas ",
	" sparkling ciders ",
	" sparkling sugars ",
	" sparkling waters ",
	" sparkling whites ",
	" spinach linguine ",
	" spinach souffles ",
	" spinach tortilla ",
	" splash grenadine ",
	" splenda granular ",
	" square focaccium ",
	" stewing chickens ",
	" strawberry juice ",
	" strawberry puree ",
	" strawberry sauce ",
	" strawberry syrup ",
	" sugar substitute ",
	" sundried tomatos ",
	" sunflower sprout ",
	" superfine sugars ",
	" swordfish steaks ",
	" szechwan peppers ",
	" taleggio cheeses ",
	" tandoori masalas ",
	" tangerine juices ",
	" tarragon parsley ",
	" tarragon vinegar ",
	" tasmanian peppers ",
	" tenderloin roast ",
	" thick asparaguss ",
	" togarashi pepper ",
	" tomatillo salsas ",
	" tomatillo sauces ",
	" tomato bouillons ",
	" tomato passatums ",
	" tomato tapenades ",
	" turbinado sugars ",
	" turkey drippings ",
	" turkey kielbasas ",
	" turkey pepperoni ",
	" turkey stuffings ",
	" turkish apricots ",
	" unbleached flour ",
	" unprocessed bran ",
	" unrefined sugars ",
	" unripe plantains ",
	" unseasoned rices ",
	" valencium orange ",
	" vanilla custards ",
	" vanilla essences ",
	" vanilla frosting ",
	" vanilla puddings ",
	" veal scallopinis ",
	" veal scaloppines ",
	" vegan chocolates ",
	" vegan margarines ",
	" vegan mayonnaise ",
	" vegetable broths ",
	" vegetable juices ",
	" vegetable sprays ",
	" vegetable stocks ",
	" vegetarian bacon ",
	" vegetarian chili ",
	" velveetum cheese ",
	" water chestnutss ",
	" watermelon juice ",
	" wheat breadcrumb ",
	" wheat fettuccine ",
	" wheat spaghettis ",
	" white asparaguss ",
	" white chocolates ",
	" white grapefruit ",
	" white peppercorn ",
	" white peppercorns ",
	" wholemeal flours ",
	" wildflower honey ",
	" winter vegetable ",
	" worchestershires ",
	" yellow cake mixs ",
	" yellow capsicums ",
	" yellow cornmeals ",
	" yellow plantains ",
	" yellow zucchinis ",
	" zucchini blossom ",
	" zucchini flowers ",
	" zucchini noodles ",
	" alfalfa sprouts ",
	" alligator meats ",
	" alligator pepper ",
	" allspice berrys ",
	" almond essences ",
	" almond extracts ",
	" almond liqueurs ",
	" alphonso mangos ",
	" amaranth flours ",
	" amarena cherrys ",
	" amaretti cookie ",
	" american cheese ",
	" anaheim peppers ",
	" anned chickpeas ",
	" apricot brandys ",
	" apricot gelatin ",
	" apricot halveds ",
	" apricot liqueur ",
	" apricot nectars ",
	" armagnac cognac ",
	" aromatic gingers ",
	" arrowroot flour ",
	" artichoke heart ",
	" asadero cheeses ",
	" asian eggplants ",
	" assorted fruits ",
	" assorted greens ",
	" assorted olives ",
	" atlantic salmon ",
	" baby artichokes ",
	" baby asparaguss ",
	" bacon drippings ",
	" bacon pancettum ",
	" banana liqueurs ",
	" banyul vinegars ",
	" barbecue sauces ",
	" barbecued porks ",
	" barbeque sauces ",
	" bechamel sauces ",
	" beef tenderloin ",
	" belgian endifes ",
	" belgian endives ",
	" black peppercorn ",
	" blackberry jams ",
	" blanc vermouths ",
	" blanched almond ",
	" blanco tequilas ",
	" bloody mary mix ",
	" blueberry juice ",
	" boston lettuces ",
	" bottle ketchups ",
	" bourbon brandys ",
	" bourbon whiskey ",
	" boursin cheeses ",
	" braeburn apples ",
	" braising greens ",
	" bratwurst links ",
	" brazilian pepper ",
	" bread stuffings ",
	" brewed espresso ",
	" broccoli crowns ",
	" broiler chicken ",
	" brown flaxseeds ",
	" browned butters ",
	" brussel sprouts ",
	" brussels sprouts ",
	" brut champagnes ",
	" buckwheat flour ",
	" buckwheat groat ",
	" buckwheat honey ",
	" buckwheat sobas ",
	" budweiser beers ",
	" bunium persicums ",
	" burratum cheese ",
	" butter lettuces ",
	" butter optional ",
	" buttered noodle ",
	" buttered toasts ",
	" butternut squash ",
	" buttery cracker ",
	" button mushroom ",
	" cabbage kimchis ",
	" cacao chocolate ",
	" caesar dressing ",
	" calamatum olive ",
	" canadian bacons ",
	" candied cherrys ",
	" candied citrons ",
	" candied gingers ",
	" candied walnuts ",
	" cannellini bean ",
	" caramel topping ",
	" carnaroli rices ",
	" cayenne peppers ",
	" cerignola olive ",
	" challah brioche ",
	" champagne grape ",
	" champagne yeast ",
	" charcoal grills ",
	" charentai melon ",
	" chayote squashs ",
	" cheddar cheeses ",
	" cheese crackers ",
	" cherry gelatins ",
	" cherry heerings ",
	" cherry liqueurs ",
	" cherry tomatoes ",
	" chervil parsley ",
	" chestnut creams ",
	" chestnut flours ",
	" chestnut honeys ",
	" chestnut purees ",
	" chicken breasts ",
	" chicken breasts ",
	" chicken carcass ",
	" chicken giblets ",
	" chicken gizzard ",
	" chicken nuggets ",
	" chicken sausage ",
	" chicken tenders ",
	" chicken thighss ",
	" chickpea flours ",
	" chickpea rinsed ",
	" chicory lettuce ",
	" chilled seltzer ",
	" chinese cabbage ",
	" chinese celerys ",
	" chinese mustard ",
	" chinese sausage ",
	" chipotle chiles ",
	" chipotle chilis ",
	" chipotle pepper ",
	" chipotle purees ",
	" chipotle salsas ",
	" chipotle sauces ",
	" chocolate cakes ",
	" chocolate chips ",
	" chocolate crust ",
	" chocolate curls ",
	" chocolate kisss ",
	" chocolate milks ",
	" chocolate sauce ",
	" chocolate syrup ",
	" chocolate wafer ",
	" chorizo sausage ",
	" ciabattum bread ",
	" ciabattum loafs ",
	" ciabattum rolls ",
	" cilantro basils ",
	" cilantro creams ",
	" cilantro leaves ",
	" cilantro sauces ",
	" cinnamon creams ",
	" cinnamon sugars ",
	" cipolline onion ",
	" cipollini onion ",
	" cocktail franks ",
	" cocktail olives ",
	" cocktail onions ",
	" cocktail peanut ",
	" cocktail sauces ",
	" coconut butters ",
	" coconut shaving ",
	" coconut sorbets ",
	" coconut vinegar ",
	" coconut yogurts ",
	" coffee liqueurs ",
	" common crackers ",
	" common mushroom ",
	" condensed milks ",
	" converted rices ",
	" coriander seeds ",
	" corn tortillass ",
	" cornbread crumb ",
	" cornflake crumb ",
	" cornmeal crusts ",
	" cortland apples ",
	" cottage cheeses ",
	" country sausage ",
	" cranberry beans ",
	" cranberry jelly ",
	" cranberry juice ",
	" cranberry sauce ",
	" cranberry vodka ",
	" cream of tartar ",
	" crema mexicanas ",
	" creme anglaises ",
	" creole mustards ",
	" crispy shallots ",
	" croissant dough ",
	" crusty baguette ",
	" cucumber salads ",
	" cultured butter ",
	" dandelion green ",
	" dark chocolates ",
	" demerara sugars ",
	" demerara syrups ",
	" desired topping ",
	" dry white wines ",
	" dulce de leches ",
	" dungeness crabs ",
	" edible glitters ",
	" egg fettuccines ",
	" egg substitutes ",
	" elbow macaronis ",
	" elephant garlic ",
	" emmental cheese ",
	" empanada doughs ",
	" enchilada sauce ",
	" english muffins ",
	" english mustard ",
	" english walnuts ",
	" enoki mushrooms ",
	" enriched flours ",
	" espresso coffee ",
	" evaporated milk ",
	" fat buttermilks ",
	" fat mayonnaises ",
	" fennel sausages ",
	" fenugreek seeds ",
	" filtered waters ",
	" finishing salts ",
	" flageolet beans ",
	" flatiron steaks ",
	" flavorless oils ",
	" florida avocado ",
	" flour tortillas ",
	" flour tortillas ",
	" focaccium bread ",
	" focaccium dough ",
	" fontina cheeses ",
	" forbidden rices ",
	" freestone peach ",
	" french baguette ",
	" french mustards ",
	" frying chickens ",
	" fuyu persimmons ",
	" garbanzo flours ",
	" garlic croutons ",
	" garlic granules ",
	" garlic sausages ",
	" gelatin dessert ",
	" german mustards ",
	" gherkin pickles ",
	" ginger liqueurs ",
	" globe artichoke ",
	" globe eggplants ",
	" golden flaxseed ",
	" graham crackers ",
	" graham crackers ",
	" grains of selims ",
	" granola cereals ",
	" grapefruit peel ",
	" greater galangal ",
	" green asparagus ",
	" green capsicums ",
	" green cardamoms ",
	" green coriander ",
	" green peppercorn ",
	" green plantains ",
	" green scallions ",
	" green zucchinis ",
	" grenadine syrup ",
	" gruyere cheeses ",
	" guajillo chiles ",
	" guajillo chilis ",
	" guinness stouts ",
	" habanero chiles ",
	" habanero chilis ",
	" habanero pepper ",
	" habanero sauces ",
	" halloumi cheese ",
	" haloumi cheeses ",
	" hamburger meats ",
	" hamburger rolls ",
	" hardboiled eggs ",
	" havarti cheeses ",
	" hazelnut butter ",
	" hazelnut flours ",
	" hazelnut spread ",
	" hazelnut syrups ",
	" hefeweizen beer ",
	" herbed croutons ",
	" herbs and spices ",
	" hibiscu flowers ",
	" honeydew melons ",
	" iceberg lettuce ",
	" iceburg lettuce ",
	" imported olives ",
	" indian bay leafs ",
	" instant coffees ",
	" instant couscou ",
	" instant tapioca ",
	" israeli couscou ",
	" italian fontina ",
	" italian parsley ",
	" italian sausage ",
	" italian tomatos ",
	" jalapeno chiles ",
	" jalapeno chilis ",
	" jalapeno pepper ",
	" jonagold apples ",
	" kabocha squashs ",
	" kalamatum olive ",
	" kasseri cheeses ",
	" kitchen bouquet ",
	" lamb tenderloin ",
	" lasagna noodles ",
	" lasagna noodles ",
	" lavender flower ",
	" lavender honeys ",
	" lemon marmalade ",
	" lesser galangals ",
	" lettuce tomatos ",
	" lime mayonnaise ",
	" lingonberry jam ",
	" littleneck clam ",
	" loaf ciabattums ",
	" lumpium wrapper ",
	" luxardo cherrys ",
	" macadamium nuts ",
	" macaroni noodle ",
	" macintosh apple ",
	" manchego cheese ",
	" mandarin juices ",
	" mandarin orange ",
	" mandarin vodkas ",
	" marcona almonds ",
	" margaritum mixs ",
	" margaritum salt ",
	" marinara sauces ",
	" marshmallow crs ",
	" mcintosh apples ",
	" merguez sausage ",
	" mexican chorizo ",
	" mexican oregano ",
	" mexican tomatos ",
	" mezzi rigatonis ",
	" milk chocolates ",
	" milk mozzarella ",
	" miniature bagel ",
	" minute tapiocas ",
	" mixed mushrooms ",
	" mixed vegetable ",
	" morel mushrooms ",
	" muenster cheese ",
	" muscovado sugar ",
	" mushroom broths ",
	" mushroom sauces ",
	" mzarella cheese ",
	" natural almonds ",
	" natural yogurts ",
	" nonpareil caper ",
	" nonstick sprays ",
	" oatmeal cookies ",
	" olive tapenades ",
	" onion scallions ",
	" orange blossoms ",
	" orange curacaos ",
	" orange gelatins ",
	" orange liqueurs ",
	" oyster crackers ",
	" oyster mushroom ",
	" oyster shuckeds ",
	" package tempehs ",
	" pancake batters ",
	" pancettum bacon ",
	" parboiled rices ",
	" parmesan cheese ",
	" parmesan shaved ",
	" parsley chervil ",
	" parsley flakess ",
	" parsley leavess ",
	" pasteurized egg ",
	" pattypan squash ",
	" peach nectarine ",
	" peanut brittles ",
	" pearled barleys ",
	" pecorino cheese ",
	" pecorino romano ",
	" peppadew pepper ",
	" pepper vinegars ",
	" peppered bacons ",
	" peruvian peppers ",
	" pickled garlics ",
	" pickled gingers ",
	" pickled herring ",
	" pickled mustard ",
	" pickled peppers ",
	" pickling spices ",
	" piloncillo cone ",
	" pimento peppers ",
	" pimiento pepper ",
	" pineapple juice ",
	" pineapple salsa ",
	" piquillo pepper ",
	" pistachio pecan ",
	" poached salmons ",
	" poblano peppers ",
	" polish kielbasa ",
	" polish sausages ",
	" polska kielbasa ",
	" pomegranate seed ",
	" popcorn kernels ",
	" popped popcorns ",
	" pork tenderloin ",
	" portobello caps ",
	" potato gnocchis ",
	" powdered sugars ",
	" premium tequila ",
	" pressure cooker ",
	" prosciutto hams ",
	" pumpkin butters ",
	" purified waters ",
	" purple cabbages ",
	" raclette cheese ",
	" rainbow quinoas ",
	" raisin currants ",
	" ranch dressings ",
	" raspberry couli ",
	" raspberry juice ",
	" raspberry puree ",
	" raspberry sauce ",
	" raspberry vodka ",
	" realemon juices ",
	" red bell pepper ",
	" remoulade sauce ",
	" rice vermicelli ",
	" ricotta cheeses ",
	" robiola cheeses ",
	" romaine lettuce ",
	" rosemary branch ",
	" rosemary thymes ",
	" rubber spatulas ",
	" ruby grapefruit ",
	" saffron strands ",
	" saffron threads ",
	" salad dressings ",
	" salt and pepper ",
	" saltine cracker ",
	" sandwich breads ",
	" saskatoon berry ",
	" satsuma oranges ",
	" sausage casings ",
	" sauvignon blanc ",
	" scallion greens ",
	" scallion whites ",
	" scotch whiskeys ",
	" seasoning salts ",
	" seltzer chilled ",
	" semolina flours ",
	" serrano peppers ",
	" seville oranges ",
	" shallot halveds ",
	" shaved parmesan ",
	" sherry vinegars ",
	" shishito pepper ",
	" shrimp bouillon ",
	" shucked oysters ",
	" sichuan peppers ",
	" silver tequilas ",
	" simmered pintos ",
	" smoked bluefish ",
	" smoked kielbasa ",
	" smoked mackerel ",
	" smoked paprikas ",
	" smoked sausages ",
	" sockeye salmons ",
	" sourdough bread ",
	" sourdough loafs ",
	" sourdough rolls ",
	" spaetzle noodle ",
	" spaghetti squash ",
	" spanish brandys ",
	" spanish chorizo ",
	" spanish paprika ",
	" spanish peanuts ",
	" sparkling cider ",
	" sparkling sugar ",
	" sparkling water ",
	" sparkling white ",
	" sparkling wines ",
	" spinach souffle ",
	" split chickpeas ",
	" squash blossoms ",
	" sriracha sauces ",
	" stale baguettes ",
	" starchy potatos ",
	" stewing chicken ",
	" stilton cheeses ",
	" stirred tahinis ",
	" strawberry jams ",
	" submarine rolls ",
	" sugar snap peas ",
	" sultana raisins ",
	" sundried tomato ",
	" sunflower seeds ",
	" superfine sugar ",
	" sweet potatoess ",
	" swordfish steak ",
	" szechwan pepper ",
	" taleggio cheese ",
	" tamarind juices ",
	" tandoori masala ",
	" tangerine juice ",
	" tangerine peels ",
	" tapioca starchs ",
	" tasmanian pepper ",
	" tender lettuces ",
	" tequila blancos ",
	" teriyaki sauces ",
	" thick asparagus ",
	" thyme rosemarys ",
	" tomatillo salsa ",
	" tomatillo sauce ",
	" tomato bouillon ",
	" tomato chutneys ",
	" tomato ketchups ",
	" tomato passatum ",
	" tomato tapenade ",
	" tonkatsu sauces ",
	" tortilla flours ",
	" tropical fruits ",
	" truffle butters ",
	" turbinado sugar ",
	" turkey carcasss ",
	" turkey dripping ",
	" turkey kielbasa ",
	" turkey sausages ",
	" turkey stuffing ",
	" turkish apricot ",
	" tzatziki sauces ",
	" unrefined sugar ",
	" unripe plantain ",
	" unseasoned rice ",
	" unsmoked bacons ",
	" vanilla custard ",
	" vanilla essence ",
	" vanilla pudding ",
	" vanilla yogurts ",
	" veal scallopini ",
	" veal scaloppine ",
	" vegan chocolate ",
	" vegan margarine ",
	" vegetable broth ",
	" vegetable juice ",
	" vegetable soups ",
	" vegetable spray ",
	" vegetable stock ",
	" vidalium onions ",
	" vienna sausages ",
	" wakame seaweeds ",
	" water chestnuts ",
	" water chestnuts ",
	" wheat baguettes ",
	" wheat farfalles ",
	" wheat macaronis ",
	" wheat rigatonis ",
	" wheat spaghetti ",
	" wheat tortillas ",
	" whipping creams ",
	" white asparagus ",
	" white chocolate ",
	" white cornmeals ",
	" white mushrooms ",
	" white peppercorn ",
	" white vermouths ",
	" wholemeal flour ",
	" wild asparaguss ",
	" wild blueberrys ",
	" wonton wrappers ",
	" worcestershires ",
	" worchestershire ",
	" yellow cake mix ",
	" yellow capsicum ",
	" yellow cornmeal ",
	" yellow mustards ",
	" yellow plantain ",
	" yellow zucchini ",
	" yellowfin tunas ",
	" young pecorinos ",
	" zucchini flower ",
	" zucchini noodle ",
	" absolut vodkas ",
	" acacium honeys ",
	" albacore tunas ",
	" aleppo peppers ",
	" alfalfa sprout ",
	" alfalfa sprouts ",
	" alfredo sauces ",
	" alligator meat ",
	" allspice berry ",
	" allspice drams ",
	" almond butters ",
	" almond essence ",
	" almond extract ",
	" almond liqueur ",
	" almond slivers ",
	" alphonso mango ",
	" amaranth flour ",
	" amaranth seeds ",
	" amarena cherry ",
	" anaheim chiles ",
	" anaheim chilis ",
	" anaheim pepper ",
	" anchovy filets ",
	" anise liqueurs ",
	" anned chickpea ",
	" apple chutneys ",
	" apple schnapps ",
	" apple vinegars ",
	" apricot brandy ",
	" apricot halved ",
	" apricot nectar ",
	" arame seaweeds ",
	" aromatic ginger ",
	" arugula pestos ",
	" arugula salads ",
	" asadero cheese ",
	" asiago cheeses ",
	" asian eggplant ",
	" assorted fruit ",
	" assorted green ",
	" assorted herbs ",
	" assorted olive ",
	" baby artichoke ",
	" baby asparagus ",
	" baby broccolis ",
	" baby cucumbers ",
	" baby eggplants ",
	" baby zucchinis ",
	" bacardi limons ",
	" bacon dripping ",
	" bag cranberrys ",
	" baking potatos ",
	" baking powders ",
	" bamboo skewers ",
	" banana liqueur ",
	" banana peppers ",
	" banyul vinegar ",
	" barbecue sauce ",
	" barbecued pork ",
	" barbeque sauce ",
	" basil parsleys ",
	" bechamel sauce ",
	" beef bouillons ",
	" beef consommes ",
	" beef drippings ",
	" beefeater gins ",
	" belgian endife ",
	" belgian endive ",
	" beluga lentils ",
	" black cardamoms ",
	" black-eyed peas ",
	" blackberry jam ",
	" blanc vermouth ",
	" blanco tequila ",
	" blood sausages ",
	" blueberry jams ",
	" borlotti beans ",
	" boston lettuce ",
	" bottle aperols ",
	" bottle ketchup ",
	" bouquet garnis ",
	" bourbon brandy ",
	" bourbon sauces ",
	" boursin cheese ",
	" box raspberrys ",
	" braeburn apple ",
	" braising green ",
	" bratwurst link ",
	" bread crackers ",
	" bread puddings ",
	" bread stuffing ",
	" brewed coffees ",
	" broccoli crown ",
	" broccoli rabes ",
	" broccoli slaws ",
	" broken walnuts ",
	" brown flaxseed ",
	" brown mustards ",
	" browned butter ",
	" brussel sprout ",
	" brussels sprout ",
	" brut champagne ",
	" buckwheat soba ",
	" budweiser beer ",
	" buffalo sauces ",
	" bunium persicum ",
	" burgundy wines ",
	" butter cookies ",
	" butter lettuce ",
	" buttered toast ",
	" cabbage kimchi ",
	" canadian bacon ",
	" candied cherry ",
	" candied citron ",
	" candied fruits ",
	" candied ginger ",
	" candied pecans ",
	" candied walnut ",
	" candy coatings ",
	" caramel sauces ",
	" caramel syrups ",
	" cardamom seeds ",
	" carnaroli rice ",
	" cashew butters ",
	" cassi liqueurs ",
	" cayenne pepper ",
	" cayenne peppers ",
	" challah breads ",
	" charcoal grill ",
	" chayote squash ",
	" cheddar cheese ",
	" cheese cracker ",
	" cheese spreads ",
	" cherry brandys ",
	" cherry gelatin ",
	" cherry heering ",
	" cherry liqueur ",
	" cherry peppers ",
	" cherry tomatos ",
	" chestnut cream ",
	" chestnut flour ",
	" chestnut honey ",
	" chestnut puree ",
	" chicken breast ",
	" chicken broths ",
	" chicken giblet ",
	" chicken livers ",
	" chicken nugget ",
	" chicken salads ",
	" chicken stocks ",
	" chicken strips ",
	" chicken tender ",
	" chicken thighs ",
	" chicken thighs ",
	" chickpea flour ",
	" chilled waters ",
	" chinese celery ",
	" chinese chilis ",
	" chipotle chile ",
	" chipotle chili ",
	" chipotle puree ",
	" chipotle salsa ",
	" chipotle sauce ",
	" chive blossoms ",
	" chocolate bars ",
	" chocolate cake ",
	" chocolate chip ",
	" chocolate curl ",
	" chocolate eggs ",
	" chocolate kiss ",
	" chocolate milk ",
	" ciabattum loaf ",
	" ciabattum roll ",
	" cider vinegars ",
	" cilantro basil ",
	" cilantro cream ",
	" cilantro mints ",
	" cilantro sauce ",
	" cinnamon barks ",
	" cinnamon cream ",
	" cinnamon sugar ",
	" clam scrubbeds ",
	" clamato juices ",
	" clotted creams ",
	" cocktail frank ",
	" cocktail olive ",
	" cocktail onion ",
	" cocktail sauce ",
	" coconut butter ",
	" coconut creams ",
	" coconut fleshs ",
	" coconut flours ",
	" coconut sorbet ",
	" coconut sugars ",
	" coconut waters ",
	" coconut yogurt ",
	" coffee filters ",
	" coffee liqueur ",
	" collard greens ",
	" colored sugars ",
	" common cracker ",
	" concord grapes ",
	" condensed milk ",
	" converted rice ",
	" cooking sprays ",
	" coriander leafs ",
	" coriander seed ",
	" coriander seeds ",
	" corn tortillas ",
	" corn tortillas ",
	" cornmeal crust ",
	" cortland apple ",
	" cotija cheeses ",
	" cottage cheese ",
	" country breads ",
	" cracker crumbs ",
	" cracker toasts ",
	" cranberry bean ",
	" cranberry jams ",
	" crawfish tails ",
	" crayfish tails ",
	" cream coconuts ",
	" crema mexicana ",
	" creme anglaise ",
	" creme fraiches ",
	" creole mustard ",
	" crispy shallot ",
	" cucumber salad ",
	" cured chorizos ",
	" curly parsleys ",
	" curly spinachs ",
	" currant jellys ",
	" daikon radishs ",
	" dark chocolate ",
	" demerara sugar ",
	" demerara syrup ",
	" dessert apples ",
	" dijon mustards ",
	" dinosaur kales ",
	" dry white wine ",
	" dulce de leche ",
	" dungeness crab ",
	" edible flowers ",
	" edible glitter ",
	" egg fettuccine ",
	" egg substitute ",
	" elbow macaroni ",
	" empanada dough ",
	" english muffin ",
	" english walnut ",
	" enoki mushroom ",
	" enriched flour ",
	" espresso beans ",
	" farmer cheeses ",
	" fat buttermilk ",
	" fat mayonnaise ",
	" favorite beers ",
	" fennel pollens ",
	" fennel sausage ",
	" fenugreek seed ",
	" fernet brancas ",
	" filtered water ",
	" finishing salt ",
	" flageolet bean ",
	" flatiron steak ",
	" flavorless oil ",
	" flaxseed meals ",
	" flour tortilla ",
	" floury potatos ",
	" fontina cheese ",
	" food colorings ",
	" forbidden rice ",
	" french lentils ",
	" french mustard ",
	" fried chickens ",
	" fried shallots ",
	" frizzled leeks ",
	" fromage blancs ",
	" fruit chutneys ",
	" fruit compotes ",
	" fryer chickens ",
	" frying chicken ",
	" fuyu persimmon ",
	" garbanzo beans ",
	" garbanzo flour ",
	" garlic confits ",
	" garlic crouton ",
	" garlic granule ",
	" garlic powders ",
	" garlic sausage ",
	" garlic sprouts ",
	" gelatin sheets ",
	" german mustard ",
	" gherkin pickle ",
	" ginger cookies ",
	" ginger liqueur ",
	" globe eggplant ",
	" glucose syrups ",
	" glutinou rices ",
	" golden raisins ",
	" graham cracker ",
	" grains of selim ",
	" grand marniers ",
	" granola cereal ",
	" grapeseed oils ",
	" greek oreganos ",
	" green cabbages ",
	" green capsicum ",
	" green cardamom ",
	" green chiliess ",
	" green pepperss ",
	" green plantain ",
	" green scallion ",
	" green zucchini ",
	" gruyere cheese ",
	" guajillo chile ",
	" guajillo chili ",
	" guinness stout ",
	" gyoza wrappers ",
	" habanero chile ",
	" habanero chili ",
	" habanero sauce ",
	" half and halfs ",
	" halibut steaks ",
	" haloumi cheese ",
	" hamburger buns ",
	" hamburger meat ",
	" hamburger roll ",
	" hardboiled egg ",
	" harissa sauces ",
	" havarti cheese ",
	" hazelnut flour ",
	" hazelnut meals ",
	" hazelnut skins ",
	" hazelnut syrup ",
	" heart of palms ",
	" heart romaines ",
	" herbed crouton ",
	" herbs and spice ",
	" hibiscu flower ",
	" honey mustards ",
	" honeydew melon ",
	" hubbard squashs ",
	" imported olive ",
	" inch baguettes ",
	" inch tortillas ",
	" indian bay leaf ",
	" instant coffee ",
	" instant flours ",
	" instant yeasts ",
	" irish whiskeys ",
	" italian basils ",
	" italian breads ",
	" italian tomato ",
	" jalapeno chile ",
	" jalapeno chili ",
	" jello gelatins ",
	" jonagold apple ",
	" juniper berrys ",
	" kabocha squash ",
	" kasseri cheese ",
	" kohlrabi bulbs ",
	" kombu seaweeds ",
	" lacinato kales ",
	" lasagna noodle ",
	" lasagna sheets ",
	" lavender honey ",
	" lemon gelatins ",
	" lesser galangal ",
	" lettuce greens ",
	" lettuce tomato ",
	" linguine finis ",
	" loaf ciabattum ",
	" long eggplants ",
	" lowfat yogurts ",
	" luxardo cherry ",
	" macadamium nut ",
	" maine lobsters ",
	" mandarin juice ",
	" mandarin vodka ",
	" mango chutneys ",
	" manioc starchs ",
	" marcona almond ",
	" margaritum mix ",
	" marinara sauce ",
	" marshmallow cr ",
	" mcintosh apple ",
	" mesclun greens ",
	" metaxa brandys ",
	" mexican cremas ",
	" mexican lagers ",
	" mexican tomato ",
	" mezzi rigatoni ",
	" milk chocolate ",
	" mineral waters ",
	" minute tapioca ",
	" mixed lettuces ",
	" mixed mushroom ",
	" monterey jacks ",
	" morel mushroom ",
	" muffin batters ",
	" mushroom broth ",
	" mushroom sauce ",
	" mushroom soups ",
	" mustard greens ",
	" natural almond ",
	" natural yogurt ",
	" nicoise olives ",
	" nigella sativas ",
	" nonfat yogurts ",
	" nonstick spray ",
	" northern beans ",
	" oatmeal cookie ",
	" oaxaca cheeses ",
	" olive tapenade ",
	" onion scallion ",
	" orange blossom ",
	" orange curacao ",
	" orange gelatin ",
	" orange lentils ",
	" orange liqueur ",
	" orange peppers ",
	" orange roughys ",
	" orange sorbets ",
	" oyster cracker ",
	" oyster shucked ",
	" package tempeh ",
	" pancake batter ",
	" paneer cheeses ",
	" panela cheeses ",
	" parboiled rice ",
	" parsley basils ",
	" parsley flakes ",
	" parsley leaves ",
	" parsley pestos ",
	" parsley thymes ",
	" pasilla chiles ",
	" passion fruits ",
	" peach gelatins ",
	" peach liqueurs ",
	" peach schnapps ",
	" peanut brittle ",
	" peanut butters ",
	" pearl couscous ",
	" pearl tapiocas ",
	" pearled barley ",
	" peasant breads ",
	" penne fusillis ",
	" pepper vinegar ",
	" peppered bacon ",
	" peruvian pepper ",
	" phyllo pastrys ",
	" picante salsas ",
	" picante sauces ",
	" pickle relishs ",
	" pickled garlic ",
	" pickled ginger ",
	" pickled onions ",
	" pickled pepper ",
	" pickling limes ",
	" pickling salts ",
	" pickling spice ",
	" pimento pepper ",
	" pineapple rums ",
	" pistachio nuts ",
	" pistachio oils ",
	" plum tomatoess ",
	" poached salmon ",
	" poblano chiles ",
	" poblano chilis ",
	" poblano pepper ",
	" poire williams ",
	" polish sausage ",
	" popcorn kernel ",
	" popped popcorn ",
	" pork spareribs ",
	" portobello cap ",
	" potato gnocchi ",
	" potato starchs ",
	" powdered sugar ",
	" premium vodkas ",
	" prosciutto ham ",
	" puffed millets ",
	" pullman breads ",
	" pumpkin butter ",
	" pumpkin fleshs ",
	" pumpkin purees ",
	" pureed bananas ",
	" pureed tomatos ",
	" purified water ",
	" purple cabbage ",
	" purple potatos ",
	" quick oatmeals ",
	" radish sprouts ",
	" rainbow chards ",
	" rainbow quinoa ",
	" rainbow trouts ",
	" raisin currant ",
	" ranch dressing ",
	" raspberry jams ",
	" realemon juice ",
	" rendered ducks ",
	" rhubarb syrups ",
	" ricotta cheese ",
	" riesling wines ",
	" roast chickens ",
	" robiola cheese ",
	" rolled barleys ",
	" romaine hearts ",
	" romano cheeses ",
	" romesco sauces ",
	" rosemary thyme ",
	" rubber spatula ",
	" russet potatos ",
	" safflower oils ",
	" saffron strand ",
	" saffron thread ",
	" salad dressing ",
	" salsa picantes ",
	" sanding sugars ",
	" sandwich bread ",
	" sandwich rolls ",
	" satsuma orange ",
	" sausage casing ",
	" savoy cabbages ",
	" scallion green ",
	" scallion white ",
	" scotch bonnets ",
	" scotch whiskey ",
	" scotch whiskys ",
	" scrambled eggs ",
	" seasoning salt ",
	" seltzer waters ",
	" semolina flour ",
	" serrano chiles ",
	" serrano chilis ",
	" serrano pepper ",
	" sesame tahinis ",
	" seville orange ",
	" shallot halved ",
	" shaoxing wines ",
	" sharp cheddars ",
	" shelling beans ",
	" sherry vinegar ",
	" shucked oyster ",
	" sichuan pepper ",
	" sichuan peppers ",
	" silver tequila ",
	" simmered pinto ",
	" sirloin steaks ",
	" sirloin strips ",
	" smoked almonds ",
	" smoked paprika ",
	" smoked salmons ",
	" smoked sausage ",
	" smoked turkeys ",
	" snack crackers ",
	" sockeye salmon ",
	" sorghum flours ",
	" sorghum syrups ",
	" sourdough loaf ",
	" sourdough roll ",
	" soy margarines ",
	" spanish brandy ",
	" spanish olives ",
	" spanish onions ",
	" spanish peanut ",
	" sparkling wine ",
	" spicy mustards ",
	" spicy sausages ",
	" spiny lobsters ",
	" split chickpea ",
	" spring garlics ",
	" squash blossom ",
	" squeeze lemons ",
	" sriracha sauce ",
	" stale baguette ",
	" starchy potato ",
	" stewed tomatos ",
	" stilton cheese ",
	" stirred tahini ",
	" strawberry jam ",
	" streaky bacons ",
	" string cheeses ",
	" submarine roll ",
	" sugar pumpkins ",
	" sugar snap pea ",
	" sultana raisin ",
	" sunflower oils ",
	" sunflower seed ",
	" sweet potatoes ",
	" tabasco peppers ",
	" tabasco sauces ",
	" tamarind juice ",
	" tamarind pulps ",
	" tangerine peel ",
	" tapioca flours ",
	" tapioca pearls ",
	" tapioca starch ",
	" tender lettuce ",
	" tequila blanco ",
	" teriyaki sauce ",
	" thyme oreganos ",
	" thyme parsleys ",
	" thyme rosemary ",
	" tomato chutney ",
	" tomato ketchup ",
	" tomato relishs ",
	" tonkatsu sauce ",
	" tortilla chips ",
	" tortilla flour ",
	" tropical fruit ",
	" truffle butter ",
	" turkey carcass ",
	" turkey sausage ",
	" tzatziki sauce ",
	" unsmoked bacon ",
	" vanilla sauces ",
	" vanilla sugars ",
	" vanilla syrups ",
	" vanilla vodkas ",
	" vanilla wafers ",
	" vanilla yogurt ",
	" vegan chickens ",
	" vegetable oils ",
	" vegetable soup ",
	" veggie burgers ",
	" venison steaks ",
	" vidalium onion ",
	" vienna sausage ",
	" vinegar sauces ",
	" virginium hams ",
	" voatsiperiferys ",
	" vodka chilleds ",
	" wakame seaweed ",
	" walnut butters ",
	" water chestnut ",
	" water chestnuts ",
	" water crackers ",
	" wheat baguette ",
	" wheat couscous ",
	" wheat crackers ",
	" wheat farfalle ",
	" wheat macaroni ",
	" wheat rigatoni ",
	" wheat tortilla ",
	" whipped creams ",
	" whipping cream ",
	" white cabbages ",
	" white cheddars ",
	" white cornmeal ",
	" white creamers ",
	" white mushroom ",
	" white tequilas ",
	" white truffles ",
	" white vermouth ",
	" white vinegars ",
	" wild asparagus ",
	" wild blueberry ",
	" wild mushrooms ",
	" winter squashs ",
	" wonton wrapper ",
	" wooden skewers ",
	" worcestershire ",
	" yellow hominys ",
	" yellow lentils ",
	" yellow mustard ",
	" yellow peppers ",
	" yellow potatos ",
	" yellow squashs ",
	" yellow tomatos ",
	" yellowfin tuna ",
	" young pecorino ",
	" absolut vodka ",
	" acacium honey ",
	" achiote seeds ",
	" acorn squashs ",
	" agave nectars ",
	" aged cheddars ",
	" albacore tuna ",
	" aleppo chilis ",
	" aleppo pepper ",
	" alfalfa sprout ",
	" alfredo sauce ",
	" allspice dram ",
	" almond butter ",
	" almond flours ",
	" almond sliver ",
	" almond syrups ",
	" amaranth seed ",
	" anaheim chile ",
	" anaheim chili ",
	" anchovy filet ",
	" anise liqueur ",
	" annatto seeds ",
	" apple brandys ",
	" apple butters ",
	" apple chutney ",
	" apple schnapp ",
	" apple vinegar ",
	" arame seaweed ",
	" arborio rices ",
	" arugula pesto ",
	" arugula salad ",
	" asiago cheese ",
	" assorted herb ",
	" assorted nuts ",
	" averna amaros ",
	" baby arugulas ",
	" baby broccoli ",
	" baby cucumber ",
	" baby eggplant ",
	" baby lettuces ",
	" baby spinachs ",
	" baby zucchini ",
	" bacardi limon ",
	" bacon greases ",
	" bag cranberry ",
	" baked potatos ",
	" bakery breads ",
	" baking apples ",
	" baking potato ",
	" baking powder ",
	" baking sprays ",
	" bamboo skewer ",
	" banana breads ",
	" banana pepper ",
	" banana purees ",
	" banana squashs ",
	" barbecue rubs ",
	" barley flours ",
	" basil leavess ",
	" basil parsley ",
	" basmati rices ",
	" beef bouillon ",
	" beef briskets ",
	" beef consomms ",
	" beef consomme ",
	" beef dripping ",
	" beef sirloins ",
	" beefeater gin ",
	" beluga lentil ",
	" beurre blancs ",
	" bibb lettuces ",
	" black cardamom ",
	" black mustards ",
	" black olivess ",
	" black peppers ",
	" black-eyed pea ",
	" blood oranges ",
	" blood sausage ",
	" blue curacaos ",
	" blueberry jam ",
	" bone removeds ",
	" borlotti bean ",
	" borlotti beans ",
	" bottle aperol ",
	" bouquet garni ",
	" bourbon sauce ",
	" box raspberry ",
	" bread cracker ",
	" bread crumbss ",
	" bread pudding ",
	" brewed coffee ",
	" brioche rolls ",
	" broccoli rabe ",
	" broccoli slaw ",
	" broken pecans ",
	" broken walnut ",
	" brown butters ",
	" brown lentils ",
	" brown mustard ",
	" brown mustards ",
	" brown raisins ",
	" buffalo sauce ",
	" bulb shallots ",
	" bulgar wheats ",
	" bulgur wheats ",
	" bulk sausages ",
	" burgundy wine ",
	" butter cookie ",
	" butter pecans ",
	" butter sauces ",
	" butterscotchs ",
	" calrose rices ",
	" can chickpeas ",
	" candied fruit ",
	" candied pecan ",
	" candied peels ",
	" candy coating ",
	" canning salts ",
	" caramel sauce ",
	" caramel syrup ",
	" caraway seeds ",
	" cardamom pods ",
	" cardamom seed ",
	" carrot juices ",
	" cashew butter ",
	" cashew creams ",
	" cassi liqueur ",
	" caster sugars ",
	" cayenne pepper ",
	" celery hearts ",
	" celery powders ",
	" chaat masalas ",
	" challah bread ",
	" cheese sauces ",
	" cheese spread ",
	" cherry brandy ",
	" cherry juices ",
	" cherry pepper ",
	" cherry syrups ",
	" cherry tomato ",
	" cherry vodkas ",
	" chicken bones ",
	" chicken broth ",
	" chicken liver ",
	" chicken meats ",
	" chicken salad ",
	" chicken skins ",
	" chicken soups ",
	" chicken stock ",
	" chicken strip ",
	" chicken thigh ",
	" chicken wings ",
	" chile peppers ",
	" chili peppers ",
	" chili powders ",
	" chilled lards ",
	" chilled water ",
	" chinese chili ",
	" chinese wines ",
	" chipotle rubs ",
	" chive blossom ",
	" chocolate bar ",
	" chocolate egg ",
	" chunky salsas ",
	" cider vinegar ",
	" cilantro mint ",
	" cinnamon bark ",
	" clam chowders ",
	" clam scrubbed ",
	" clamato juice ",
	" clotted cream ",
	" clover honeys ",
	" cocoa butters ",
	" coconut chips ",
	" coconut cream ",
	" coconut flesh ",
	" coconut flour ",
	" coconut meats ",
	" coconut milks ",
	" coconut sugar ",
	" coconut water ",
	" coffee filter ",
	" colby cheeses ",
	" coleslaw mixs ",
	" collard green ",
	" collard greens ",
	" colored sugar ",
	" comte cheeses ",
	" concord grape ",
	" cookie crumbs ",
	" cooking spray ",
	" coriander leaf ",
	" coriander seed ",
	" corn kernelss ",
	" corn shuckeds ",
	" corn tortilla ",
	" cotija cheese ",
	" country bread ",
	" country loafs ",
	" cracker crumb ",
	" cracker toast ",
	" cranberry jam ",
	" crawfish tail ",
	" crayfish tail ",
	" cream cheeses ",
	" cream coconut ",
	" cream sherrys ",
	" creamed corns ",
	" creme fraiche ",
	" creole spices ",
	" crusty breads ",
	" cured chorizo ",
	" curly endives ",
	" curly parsley ",
	" curly spinach ",
	" currant jelly ",
	" curry powders ",
	" daikon radish ",
	" demerara rums ",
	" dessert apple ",
	" dijon mustard ",
	" dill parsleys ",
	" dinosaur kale ",
	" double creams ",
	" dry red wines ",
	" edamame beans ",
	" edible flower ",
	" espresso bean ",
	" farmer cheese ",
	" favorite beer ",
	" fennel fronds ",
	" fennel pollen ",
	" fernet branca ",
	" fetum cheeses ",
	" filet mignons ",
	" flat anchovys ",
	" flat parsleys ",
	" flaxseed meal ",
	" flaxseed oils ",
	" floral honeys ",
	" floury potato ",
	" food coloring ",
	" french breads ",
	" french lentil ",
	" french toasts ",
	" fresno chiles ",
	" fried chicken ",
	" fried potatos ",
	" fried shallot ",
	" frizzled leek ",
	" fromage blanc ",
	" fruit chutney ",
	" fruit compote ",
	" fryer chicken ",
	" gaetum olives ",
	" garam masalas ",
	" garbanzo bean ",
	" garlic broths ",
	" garlic chives ",
	" garlic chivess ",
	" garlic confit ",
	" garlic hummus ",
	" garlic juices ",
	" garlic powder ",
	" garlic scapes ",
	" garlic sprout ",
	" gefilte fishs ",
	" gelatin sheet ",
	" giblet broths ",
	" gigante beans ",
	" ginger cookie ",
	" ginger juices ",
	" ginger syrups ",
	" globe tomatos ",
	" glucose syrup ",
	" gluten flours ",
	" glutinou rice ",
	" gold tequilas ",
	" golden raisin ",
	" golden syrups ",
	" gouda cheeses ",
	" grana padanos ",
	" grand marnier ",
	" grape tomatos ",
	" grapeseed oil ",
	" greek oregano ",
	" greek yogurts ",
	" green bananas ",
	" green cabbage ",
	" green cherrys ",
	" green chilies ",
	" green garlics ",
	" green lentils ",
	" green onionss ",
	" green papayas ",
	" green peppers ",
	" green peppers ",
	" green tomatos ",
	" guava nectars ",
	" gyoza wrapper ",
	" half and half ",
	" halibut steak ",
	" halved lemons ",
	" hamburger bun ",
	" hanger steaks ",
	" haricot verts ",
	" harissa sauce ",
	" hazelnut meal ",
	" hazelnut oils ",
	" hazelnut skin ",
	" heart of palm ",
	" heart romaine ",
	" hearty breads ",
	" hearty greens ",
	" herb bouquets ",
	" herb mixtures ",
	" hickory chips ",
	" hoisin sauces ",
	" honey mustard ",
	" hubbard squash ",
	" inch baguette ",
	" inch tortilla ",
	" instant flour ",
	" instant grits ",
	" instant milks ",
	" instant yeast ",
	" iodized salts ",
	" irish whiskey ",
	" italian basil ",
	" italian bread ",
	" italian herbs ",
	" italian plums ",
	" italian tunas ",
	" jamaican rums ",
	" jarred salsas ",
	" jasmine rices ",
	" jello gelatin ",
	" jigger vodkas ",
	" juice oranges ",
	" jumbo shrimps ",
	" juniper berry ",
	" juniper berrys ",
	" kidney beanss ",
	" kimchi juices ",
	" kirschwassers ",
	" kohlrabi bulb ",
	" kombu seaweed ",
	" lacinato kale ",
	" lasagna sheet ",
	" laughing cows ",
	" lavender buds ",
	" lemon gelatin ",
	" lemon halveds ",
	" lemon peppers ",
	" lemon verbenas ",
	" lemon yogurts ",
	" lettuce green ",
	" lillet blancs ",
	" linguine fini ",
	" link sausages ",
	" litsea cubebas ",
	" little waters ",
	" live lobsters ",
	" loaf brioches ",
	" loaf challahs ",
	" lobster meats ",
	" lobster tails ",
	" london broils ",
	" long eggplant ",
	" lowfat yogurt ",
	" lychee juices ",
	" madeira wines ",
	" maine lobster ",
	" malt vinegars ",
	" mango chutney ",
	" mango sorbets ",
	" manioc starch ",
	" marsala wines ",
	" marshmallowss ",
	" matzo farfels ",
	" meat fillings ",
	" meatloaf mixs ",
	" medjool dates ",
	" mesclun green ",
	" metal skewers ",
	" metaxa brandy ",
	" mexican beers ",
	" mexican crema ",
	" mexican lager ",
	" mexican rices ",
	" millet flours ",
	" mineral water ",
	" mint garnishs ",
	" mint parsleys ",
	" miracle whips ",
	" mixed lettuce ",
	" mixed tomatos ",
	" monterey jack ",
	" muffin batter ",
	" mushroom caps ",
	" mushroom soup ",
	" mustard green ",
	" mustard greens ",
	" mustard seeds ",
	" napa cabbages ",
	" navel oranges ",
	" nicoise olive ",
	" nigella sativa ",
	" nigella seeds ",
	" nonfat yogurt ",
	" nori seaweeds ",
	" northern bean ",
	" oaxaca cheese ",
	" onion powders ",
	" orange juices ",
	" orange lentil ",
	" orange pepper ",
	" orange roughy ",
	" orange sauces ",
	" orange sorbet ",
	" orange wheels ",
	" orgeat syrups ",
	" oyster sauces ",
	" oz spaghettis ",
	" pan drippings ",
	" panch phorons ",
	" paneer cheese ",
	" panela cheese ",
	" parsley basil ",
	" parsley dills ",
	" parsley mints ",
	" parsley pesto ",
	" parsley thyme ",
	" pasilla chile ",
	" passion fruit ",
	" pastry brushs ",
	" pastry creams ",
	" pastry doughs ",
	" pastry flours ",
	" peach brandys ",
	" peach gelatin ",
	" peach liqueur ",
	" peach nectars ",
	" peach schnapp ",
	" peach yogurts ",
	" peanut butter ",
	" peanut sauces ",
	" pear liqueurs ",
	" pearl barleys ",
	" pearl couscou ",
	" pearl tapioca ",
	" peasant bread ",
	" pecan walnuts ",
	" penne fusilli ",
	" penne rigates ",
	" pepper jellys ",
	" pepper sauces ",
	" pepper tastes ",
	" pepper vodkas ",
	" pepperoncinis ",
	" pequin chiles ",
	" percent milks ",
	" petrale soles ",
	" phyllo doughs ",
	" phyllo pastry ",
	" phyllo sheets ",
	" picante salsa ",
	" picante sauce ",
	" pickle brines ",
	" pickle juices ",
	" pickle relish ",
	" pickled beets ",
	" pickled onion ",
	" pickled pears ",
	" pickling lime ",
	" pickling salt ",
	" pineapple rum ",
	" pippin apples ",
	" pistachio nut ",
	" pistachio oil ",
	" pitum pockets ",
	" plastic wraps ",
	" plum tomatoes ",
	" plymouth gins ",
	" poblano chile ",
	" poblano chili ",
	" poire william ",
	" pork sausages ",
	" pork sparerib ",
	" porridge oats ",
	" potato breads ",
	" potato flours ",
	" potato salads ",
	" potato starch ",
	" premium vodka ",
	" pretzel salts ",
	" prickly pears ",
	" puffed millet ",
	" pullman bread ",
	" pumpernickels ",
	" pumpkin flesh ",
	" pumpkin puree ",
	" pumpkin seeds ",
	" pureed banana ",
	" pureed tomato ",
	" purple basils ",
	" purple onions ",
	" purple potato ",
	" quark cheeses ",
	" queso blancos ",
	" queso frescos ",
	" quick oatmeal ",
	" quinoa flours ",
	" radish greens ",
	" radish sprout ",
	" rainbow chard ",
	" rainbow trout ",
	" raisin breads ",
	" raman noodles ",
	" rapeseed oils ",
	" raspberry jam ",
	" real vanillas ",
	" refried beans ",
	" rendered duck ",
	" rhubarb syrup ",
	" rice crackers ",
	" rice vinegars ",
	" riesling wine ",
	" risotto rices ",
	" ritz crackers ",
	" roast chicken ",
	" roast turkeys ",
	" roasting pans ",
	" roll wrappers ",
	" rolled barley ",
	" romaine heart ",
	" romano cheese ",
	" romesco sauce ",
	" rose essences ",
	" rotel tomatos ",
	" russet potato ",
	" safflower oil ",
	" sage sausages ",
	" salmon filets ",
	" salmon steaks ",
	" salsa frescas ",
	" salsa picante ",
	" sambal oeleks ",
	" sanding sugar ",
	" sandwich buns ",
	" sandwich roll ",
	" sausage links ",
	" sausage meats ",
	" savoy cabbage ",
	" scalded milks ",
	" scotch bonnet ",
	" scotch whisky ",
	" scrambled egg ",
	" season blends ",
	" seltzer water ",
	" serrano chile ",
	" serrano chili ",
	" sesame seedss ",
	" sesame tahini ",
	" shaoxing wine ",
	" sharp cheddar ",
	" shelling bean ",
	" sichuan pepper ",
	" sifted flours ",
	" single creams ",
	" sirloin beefs ",
	" sirloin steak ",
	" sirloin strip ",
	" skimmed milks ",
	" smoked almond ",
	" smoked bacons ",
	" smoked goudas ",
	" smoked salmon ",
	" smoked trouts ",
	" smoked turkey ",
	" snack cracker ",
	" soda crackers ",
	" soft caramels ",
	" soppressatums ",
	" sorghum flour ",
	" sorghum syrup ",
	" soy margarine ",
	" spanish olive ",
	" spanish onion ",
	" spanish rices ",
	" spicy mustard ",
	" spicy sausage ",
	" spiny lobster ",
	" splash waters ",
	" spring garlic ",
	" spring greens ",
	" spring onions ",
	" spring waters ",
	" squeeze lemon ",
	" squeeze limes ",
	" steamed rices ",
	" stewed tomato ",
	" stewing beefs ",
	" stewing lambs ",
	" strawberriess ",
	" streaky bacon ",
	" string cheese ",
	" striped basss ",
	" stuffing mixs ",
	" sturdy greens ",
	" sugar cookies ",
	" sugar neededs ",
	" sugar pumpkin ",
	" sunflower oil ",
	" swiss cheeses ",
	" tabasco pepper ",
	" tabasco sauce ",
	" tahini sauces ",
	" tamari sauces ",
	" tamarind pulp ",
	" tapioca flour ",
	" tapioca pearl ",
	" tartar sauces ",
	" tempeh bacons ",
	" thick yogurts ",
	" thyme branchs ",
	" thyme leavess ",
	" thyme oregano ",
	" thyme parsley ",
	" tiger shrimps ",
	" tofu crumbles ",
	" tomato basils ",
	" tomato juices ",
	" tomato pastes ",
	" tomato pestos ",
	" tomato purees ",
	" tomato relish ",
	" tomato salsas ",
	" tomato sauces ",
	" tortilla chip ",
	" truffle salts ",
	" turkey bacons ",
	" turkey broths ",
	" turkey gravys ",
	" turkey thighs ",
	" turnip greens ",
	" vanilla beans ",
	" vanilla sauce ",
	" vanilla sugar ",
	" vanilla syrup ",
	" vanilla vodka ",
	" vanilla wafer ",
	" veal scallops ",
	" vegan butters ",
	" vegan cheeses ",
	" vegan chicken ",
	" vegetable oil ",
	" veggie burger ",
	" venison steak ",
	" vialone nanos ",
	" vinegar sauce ",
	" virginium ham ",
	" voatsiperifery ",
	" vodka chilled ",
	" wafer cookies ",
	" walnut butter ",
	" walnut flours ",
	" walnut pecans ",
	" water chestnut ",
	" water cracker ",
	" wheat couscou ",
	" wheat cracker ",
	" wheat glutens ",
	" whipped cream ",
	" white cabbage ",
	" white cheddar ",
	" white cheeses ",
	" white creamer ",
	" white hominys ",
	" white mustards ",
	" white peppers ",
	" white potatos ",
	" white quinoas ",
	" white radishs ",
	" white raisins ",
	" white tequila ",
	" white truffle ",
	" white turnips ",
	" white vinegar ",
	" wild arugulas ",
	" wild mushroom ",
	" wine vinegars ",
	" winter squash ",
	" wondra flours ",
	" wooden skewer ",
	" yellow apples ",
	" yellow chiles ",
	" yellow hominy ",
	" yellow lentil ",
	" yellow onions ",
	" yellow pepper ",
	" yellow potato ",
	" yellow squash ",
	" yellow tomato ",
	" yogurt sauces ",
	" young carrots ",
	" achiote seed ",
	" acorn squash ",
	" acorn squashs ",
	" adobo sauces ",
	" adzuki beans ",
	" agave nectar ",
	" agave syrups ",
	" aged cheddar ",
	" ajwain seeds ",
	" alaskan cods ",
	" aleppo chili ",
	" almond flour ",
	" almond meals ",
	" almond milks ",
	" almond skins ",
	" almond syrup ",
	" amontillados ",
	" ancho chiles ",
	" ancho chilis ",
	" annatto oils ",
	" annatto seed ",
	" apple brandy ",
	" apple butter ",
	" apple ciders ",
	" apple jellys ",
	" apple juices ",
	" apple purees ",
	" apple sauces ",
	" apple syrups ",
	" apricot jams ",
	" arbol chiles ",
	" arborio rice ",
	" arctic chars ",
	" arepa flours ",
	" assorted nut ",
	" averna amaro ",
	" avocado dips ",
	" avocado oils ",
	" baby arugula ",
	" baby carrots ",
	" baby fennels ",
	" baby lettuce ",
	" baby potatos ",
	" baby rockets ",
	" baby shrimps ",
	" baby spinach ",
	" baby squashs ",
	" baby tomatos ",
	" baby turnips ",
	" bacardi rums ",
	" bacon grease ",
	" bacon strips ",
	" bag spinachs ",
	" baked potato ",
	" bakery bread ",
	" baking apple ",
	" baking sodas ",
	" baking spray ",
	" banana bread ",
	" banana chips ",
	" banana puree ",
	" banana squash ",
	" barbecue rub ",
	" barley flour ",
	" barley malts ",
	" basil chives ",
	" basil leaves ",
	" basil pestos ",
	" basmati rice ",
	" bay scallops ",
	" bean sprouts ",
	" bean threads ",
	" beef brisket ",
	" beef consomm ",
	" beef marrows ",
	" beef sirloin ",
	" beef tongues ",
	" bell peppers ",
	" benedictines ",
	" besan flours ",
	" beurre blanc ",
	" bibb lettuce ",
	" bing cherrys ",
	" bird peppers ",
	" biscuit mixs ",
	" black beanss ",
	" black mustard ",
	" black olives ",
	" black pepper ",
	" blackcurrants ",
	" blended oils ",
	" bleu cheeses ",
	" blood orange ",
	" blood oranges ",
	" blue cheeses ",
	" blue curacao ",
	" blueberriess ",
	" bone removed ",
	" borlotti bean ",
	" bourbon ryes ",
	" boysenberrys ",
	" bread crumbs ",
	" bread crumbs ",
	" bread flours ",
	" breadcrumbss ",
	" brie cheeses ",
	" brioche buns ",
	" brioche roll ",
	" broccoflowers ",
	" broken pecan ",
	" brook trouts ",
	" brown butter ",
	" brown lentil ",
	" brown mustard ",
	" brown onions ",
	" brown raisin ",
	" brown sugars ",
	" bruschettums ",
	" bulb fennels ",
	" bulb garlics ",
	" bulb shallot ",
	" bulgar wheat ",
	" bulgur wheat ",
	" bulk sausage ",
	" butter beans ",
	" butter pecan ",
	" butter sauce ",
	" butterscotch ",
	" californiums ",
	" calrose rice ",
	" can chickpea ",
	" canary melons ",
	" candied peel ",
	" canning salt ",
	" caper berrys ",
	" caper brines ",
	" caper juices ",
	" caraway seed ",
	" cardamom pod ",
	" carne asadas ",
	" carrot curls ",
	" carrot juice ",
	" cashew cream ",
	" cashew halfs ",
	" cashew milks ",
	" caster sugar ",
	" cauliflowers ",
	" celery heart ",
	" celery powder ",
	" celery salts ",
	" celery seeds ",
	" chaat masala ",
	" chanterelles ",
	" chat masalas ",
	" cheese curds ",
	" cheese sauce ",
	" cheese soups ",
	" cheese whizs ",
	" cheesecloths ",
	" cherry colas ",
	" cherry juice ",
	" cherry syrup ",
	" cherry vodka ",
	" chex cereals ",
	" chicken bone ",
	" chicken fats ",
	" chicken meat ",
	" chicken skin ",
	" chicken soup ",
	" chicken wing ",
	" chile pepper ",
	" chile salsas ",
	" chile sauces ",
	" chili pepper ",
	" chili peppers ",
	" chili powder ",
	" chili sauces ",
	" chilled lard ",
	" chimichurris ",
	" chinese wine ",
	" chipotle rub ",
	" chuck roasts ",
	" chuck steaks ",
	" chunky salsa ",
	" citric acids ",
	" citru fruits ",
	" citru juices ",
	" citru vodkas ",
	" clam chowder ",
	" clear honeys ",
	" clover honey ",
	" cocoa butter ",
	" coconut chip ",
	" coconut meat ",
	" coconut milk ",
	" coconut oils ",
	" coconut rums ",
	" coffee beans ",
	" coffee cakes ",
	" coho salmons ",
	" colby cheese ",
	" coleslaw mix ",
	" collard green ",
	" comice pears ",
	" comt cheeses ",
	" comte cheese ",
	" cookie crumb ",
	" cooking oils ",
	" corn kernels ",
	" corn kernels ",
	" corn niblets ",
	" corn relishs ",
	" corn shucked ",
	" corn starchs ",
	" corned beefs ",
	" cornish hens ",
	" corona beers ",
	" country hams ",
	" country loaf ",
	" cranberriess ",
	" cream cheese ",
	" cream sauces ",
	" cream sherry ",
	" creamed corn ",
	" creole spice ",
	" crisp apples ",
	" crisp bacons ",
	" crme fraches ",
	" crusty bread ",
	" crusty rolls ",
	" curing salts ",
	" curly endive ",
	" curry powder ",
	" custard mixs ",
	" cut chickens ",
	" deli turkeys ",
	" demerara rum ",
	" dill parsley ",
	" dill pickles ",
	" dinner rolls ",
	" double cream ",
	" dry mustards ",
	" dry red wine ",
	" duck confits ",
	" edam cheeses ",
	" edamame bean ",
	" egg noodless ",
	" egg whiskeds ",
	" english peas ",
	" fennel bulbs ",
	" fennel frond ",
	" fennel seeds ",
	" feta cheeses ",
	" fetum cheese ",
	" filet mignon ",
	" filo pastrys ",
	" fino sherrys ",
	" flank steaks ",
	" flat anchovy ",
	" flat parsley ",
	" flaxseed oil ",
	" floral honey ",
	" frankfurters ",
	" french beans ",
	" french bread ",
	" french rolls ",
	" french toast ",
	" fresno chile ",
	" fried onions ",
	" fried potato ",
	" fruit juices ",
	" fruit salads ",
	" fruit salsas ",
	" gaetum olive ",
	" garam masala ",
	" garlic broth ",
	" garlic bulbs ",
	" garlic chive ",
	" garlic chives ",
	" garlic hummu ",
	" garlic juice ",
	" garlic salts ",
	" garlic scape ",
	" gefilte fish ",
	" gem lettuces ",
	" giblet broth ",
	" gigante bean ",
	" ginger beers ",
	" ginger juice ",
	" ginger snaps ",
	" ginger syrup ",
	" gingerbreads ",
	" globe tomato ",
	" gluten flour ",
	" gluten frees ",
	" goat cheeses ",
	" gold tequila ",
	" golden beets ",
	" golden syrup ",
	" gouda cheese ",
	" grana padano ",
	" granola bars ",
	" grape jellys ",
	" grape tomato ",
	" greek fetums ",
	" greek olives ",
	" greek yogurt ",
	" green apples ",
	" green banana ",
	" green beanss ",
	" green cherry ",
	" green chiles ",
	" green chilis ",
	" green garlic ",
	" green grapes ",
	" green lentil ",
	" green mangos ",
	" green olives ",
	" green onions ",
	" green onions ",
	" green papaya ",
	" green pepper ",
	" green peppers ",
	" green salsas ",
	" green tomato ",
	" guava juices ",
	" guava nectar ",
	" haa avocados ",
	" halved lemon ",
	" hanger steak ",
	" hardy greens ",
	" haricot vert ",
	" hazelnut oil ",
	" hearty bread ",
	" hearty green ",
	" heavy creams ",
	" herb bouquet ",
	" herb butters ",
	" herb mixture ",
	" hickory chip ",
	" hoisin sauce ",
	" hollandaises ",
	" hominy grits ",
	" honey syrups ",
	" horseradishs ",
	" huckleberrys ",
	" icing sugars ",
	" inch celerys ",
	" inch gingers ",
	" instant grit ",
	" instant milk ",
	" iodized salt ",
	" irish creams ",
	" italian herb ",
	" italian plum ",
	" italian tuna ",
	" jack cheeses ",
	" jamaica rums ",
	" jamaican rum ",
	" jarred salsa ",
	" jasmine rice ",
	" jigger vodka ",
	" juice orange ",
	" jumbo shrimp ",
	" juniper berry ",
	" kaffir limes ",
	" kaiser rolls ",
	" kernel corns ",
	" ketjap manis ",
	" kidney beans ",
	" kidney beans ",
	" kimchi juice ",
	" king salmons ",
	" kirschwasser ",
	" kosher salts ",
	" langoustines ",
	" laughing cow ",
	" lavender bud ",
	" leafy greens ",
	" lemon aiolis ",
	" lemon basils ",
	" lemon grasss ",
	" lemon halved ",
	" lemon juices ",
	" lemon myrtles ",
	" lemon pepper ",
	" lemon thymes ",
	" lemon verbena ",
	" lemon vodkas ",
	" lemon wheels ",
	" lemon yogurt ",
	" lillet blanc ",
	" lime butters ",
	" lime halveds ",
	" lingonberrys ",
	" link sausage ",
	" litsea cubeba ",
	" little water ",
	" live lobster ",
	" live mussels ",
	" loaf brioche ",
	" loaf challah ",
	" lobster meat ",
	" lobster tail ",
	" london broil ",
	" lowfat milks ",
	" lychee juice ",
	" madeira wine ",
	" madra currys ",
	" maldon salts ",
	" malt vinegar ",
	" mango purees ",
	" mango salsas ",
	" mango sorbet ",
	" mango-gingers ",
	" manila clams ",
	" maple sugars ",
	" maple syrups ",
	" mara peppers ",
	" marrow bones ",
	" marsala wine ",
	" marshmallows ",
	" marshmallows ",
	" masa harinas ",
	" matzo farfel ",
	" matzoh meals ",
	" meat filling ",
	" meatloaf mix ",
	" medjool date ",
	" metal skewer ",
	" mexican beer ",
	" mexican rice ",
	" meyer lemons ",
	" milk yogurts ",
	" millet flour ",
	" mint garnish ",
	" mint leavess ",
	" mint parsley ",
	" minute rices ",
	" miracle whip ",
	" mixed berrys ",
	" mixed citrus ",
	" mixed fruits ",
	" mixed greens ",
	" mixed olives ",
	" mixed spices ",
	" mixed tomato ",
	" muscat wines ",
	" mushroom cap ",
	" mustard green ",
	" mustard oils ",
	" mustard seed ",
	" napa cabbage ",
	" navel orange ",
	" neutral oils ",
	" nigella seed ",
	" nilla wafers ",
	" noilly prats ",
	" nonfat milks ",
	" nori seaweed ",
	" nut biscuits ",
	" onion powder ",
	" orange juice ",
	" orange peels ",
	" orange sauce ",
	" orange sodas ",
	" orange wheel ",
	" orecchiettes ",
	" oreo cookies ",
	" orgeat syrup ",
	" oyster sauce ",
	" oz parmesans ",
	" oz proseccos ",
	" oz spaghetti ",
	" pacific cods ",
	" paella rices ",
	" pan dripping ",
	" pancake mixs ",
	" panch phoron ",
	" papaya seeds ",
	" pappardelles ",
	" parsley dill ",
	" parsley mint ",
	" passionfruits ",
	" pastry brush ",
	" pastry cream ",
	" pastry dough ",
	" pastry flour ",
	" pea tendrils ",
	" peach brandy ",
	" peach nectar ",
	" peach purees ",
	" peach syrups ",
	" peach yogurt ",
	" peanut sauce ",
	" pear brandys ",
	" pear liqueur ",
	" pear nectars ",
	" pear tomatos ",
	" pearl barley ",
	" pearl onions ",
	" pearl sugars ",
	" pecan walnut ",
	" peking ducks ",
	" penne rigate ",
	" peperoncinis ",
	" peperoncinos ",
	" pepper jelly ",
	" pepper sauce ",
	" pepper taste ",
	" pepper vodka ",
	" pepperoncini ",
	" pequin chile ",
	" percent milk ",
	" perciatellis ",
	" pesto sauces ",
	" petrale sole ",
	" phyllo dough ",
	" phyllo sheet ",
	" pickle brine ",
	" pickle chips ",
	" pickle juice ",
	" pickled beet ",
	" pickled pear ",
	" pig trotters ",
	" pignoli nuts ",
	" pippin apple ",
	" pitum breads ",
	" pitum pocket ",
	" pizza crusts ",
	" pizza doughs ",
	" pizza sauces ",
	" plastic wrap ",
	" plum tomatos ",
	" plymouth gin ",
	" poached eggs ",
	" pomegranates ",
	" ponzu sauces ",
	" popped corns ",
	" pork sausage ",
	" porridge oat ",
	" potato bread ",
	" potato chips ",
	" potato flour ",
	" potato rolls ",
	" potato salad ",
	" pretzel salt ",
	" prickly pear ",
	" prune juices ",
	" prune purees ",
	" pudding mixs ",
	" puff pastrys ",
	" puffed rices ",
	" pumpernickel ",
	" pumpkin pies ",
	" pumpkin seed ",
	" pumpkinseeds ",
	" purple basil ",
	" purple onion ",
	" purple plums ",
	" quark cheese ",
	" quatre pices ",
	" queso blanco ",
	" queso fresco ",
	" quick breads ",
	" quinoa flour ",
	" radish green ",
	" raisin bread ",
	" raman noodle ",
	" rapeseed oil ",
	" raspberriess ",
	" real butters ",
	" real vanilla ",
	" red pepperss ",
	" refried bean ",
	" rice cereals ",
	" rice cracker ",
	" rice krispys ",
	" rice noodles ",
	" rice vinegar ",
	" risotto rice ",
	" rittenhouses ",
	" ritz cracker ",
	" roast turkey ",
	" roasting pan ",
	" roll wrapper ",
	" roma tomatos ",
	" romano beans ",
	" rose essence ",
	" rotel tomato ",
	" royal icings ",
	" rubbed sages ",
	" rum liqueurs ",
	" rye whiskeys ",
	" sage sausage ",
	" salad greens ",
	" salmon filet ",
	" salmon skins ",
	" salmon steak ",
	" salsa fresca ",
	" salsa verdes ",
	" salt peppers ",
	" sambal oelek ",
	" san marzanos ",
	" sandwich bun ",
	" sarsaparillas ",
	" satay sauces ",
	" sausage link ",
	" sausage meat ",
	" scalded milk ",
	" season blend ",
	" season salts ",
	" seckel pears ",
	" serrano hams ",
	" sesame seeds ",
	" sesame seeds ",
	" shallot oils ",
	" sheet matzos ",
	" sherry wines ",
	" sifted flour ",
	" silken tofus ",
	" single cream ",
	" sirloin beef ",
	" sirloin tips ",
	" skimmed milk ",
	" skirt steaks ",
	" smith apples ",
	" smoked bacon ",
	" smoked fishs ",
	" smoked gouda ",
	" smoked salts ",
	" smoked tofus ",
	" smoked trout ",
	" smoky bacons ",
	" soba noodles ",
	" soda cracker ",
	" soft butters ",
	" soft caramel ",
	" soft cheeses ",
	" soppressatum ",
	" sour cherrys ",
	" sour pickles ",
	" soy proteins ",
	" soybean oils ",
	" spaghettinis ",
	" spanish rice ",
	" spelt berrys ",
	" spelt flours ",
	" spicy greens ",
	" splash water ",
	" sponge cakes ",
	" spring green ",
	" spring onion ",
	" spring water ",
	" squash seeds ",
	" squeeze lime ",
	" stale breads ",
	" steak sauces ",
	" steamed rice ",
	" stewing beef ",
	" stewing lamb ",
	" sticky rices ",
	" strawberries ",
	" string beans ",
	" strip bacons ",
	" strip steaks ",
	" striped bass ",
	" stuffing mix ",
	" sturdy green ",
	" sugar cookie ",
	" sugar needed ",
	" sugar syrups ",
	" sweet onions ",
	" sweet potatos ",
	" swiss chards ",
	" swiss cheese ",
	" table creams ",
	" tagliatelles ",
	" tahini sauce ",
	" tamari sauce ",
	" tart cherrys ",
	" tartar sauce ",
	" tempeh bacon ",
	" tender herbs ",
	" tepid waters ",
	" thai peppers ",
	" thick bacons ",
	" thick yogurt ",
	" thyme branch ",
	" thyme honeys ",
	" thyme leaves ",
	" tiger prawns ",
	" tiger shrimp ",
	" tiny potatos ",
	" tium mariums ",
	" toast points ",
	" tofu crumble ",
	" tomato basil ",
	" tomato juice ",
	" tomato paste ",
	" tomato pesto ",
	" tomato puree ",
	" tomato salsa ",
	" tomato sauce ",
	" tomato soups ",
	" tonic waters ",
	" tortiglionis ",
	" truffle oils ",
	" truffle salt ",
	" turkey bacon ",
	" turkey bones ",
	" turkey broth ",
	" turkey gravy ",
	" turkey meats ",
	" turkey thigh ",
	" turkey wings ",
	" turnip green ",
	" udon noodles ",
	" vanilla bean ",
	" vanilla pods ",
	" veal scallop ",
	" vegan butter ",
	" vegan cheese ",
	" vegan sugars ",
	" vialone nano ",
	" vinaigrettes ",
	" vsop cognacs ",
	" wafer cookie ",
	" walnut flour ",
	" walnut halfs ",
	" walnut pecan ",
	" waxy potatos ",
	" wheat bagels ",
	" wheat berrys ",
	" wheat breads ",
	" wheat flours ",
	" wheat gluten ",
	" wheat pennes ",
	" wheat pitums ",
	" wheat toasts ",
	" white cheese ",
	" white flours ",
	" white grapes ",
	" white hominy ",
	" white mustard ",
	" white onions ",
	" white peachs ",
	" white pepper ",
	" white potato ",
	" white quinoa ",
	" white radish ",
	" white radishs ",
	" white raisin ",
	" white sauces ",
	" white sugars ",
	" white turnip ",
	" white wheats ",
	" wild arugula ",
	" wild salmons ",
	" wine vinegar ",
	" wondra flour ",
	" wonton skins ",
	" xanthan gums ",
	" yellow apple ",
	" yellow beets ",
	" yellow chile ",
	" yellow corns ",
	" yellow finns ",
	" yellow grits ",
	" yellow misos ",
	" yellow onion ",
	" yellow rices ",
	" yogurt sauce ",
	" young carrot ",
	" acorn squash ",
	" adobo sauce ",
	" adzuki bean ",
	" agave syrup ",
	" ajwain seed ",
	" alaskan cod ",
	" almond meal ",
	" almond milk ",
	" almond oils ",
	" almond skin ",
	" amber beers ",
	" amontillado ",
	" ancho chile ",
	" ancho chili ",
	" anise seeds ",
	" anjou pears ",
	" annatto oil ",
	" apple cakes ",
	" apple cider ",
	" apple jelly ",
	" apple juice ",
	" apple pears ",
	" apple puree ",
	" apple sauce ",
	" apple syrup ",
	" applesauces ",
	" apricot jam ",
	" arbol chile ",
	" arctic char ",
	" arepa flour ",
	" asafoetidas ",
	" asian pears ",
	" avocado dip ",
	" avocado oil ",
	" baby capers ",
	" baby carrot ",
	" baby fennel ",
	" baby greens ",
	" baby onions ",
	" baby potato ",
	" baby rocket ",
	" baby shrimp ",
	" baby squash ",
	" baby tomato ",
	" baby turnip ",
	" bacardi rum ",
	" bacon strip ",
	" bag spinach ",
	" bagel chips ",
	" baked beans ",
	" baked tofus ",
	" baking mixs ",
	" baking soda ",
	" banana chip ",
	" banana rums ",
	" barley malt ",
	" basil chive ",
	" basil mints ",
	" basil pesto ",
	" bay leavess ",
	" bay scallop ",
	" bay shrimps ",
	" bean broths ",
	" bean purees ",
	" bean salads ",
	" bean sprout ",
	" bean sprouts ",
	" bean thread ",
	" beaten eggs ",
	" beef broths ",
	" beef cheeks ",
	" beef chucks ",
	" beef livers ",
	" beef marrow ",
	" beef minces ",
	" beef roasts ",
	" beef steaks ",
	" beef stocks ",
	" beef tongue ",
	" beet greens ",
	" bell pepper ",
	" bell peppers ",
	" benedictine ",
	" benne seeds ",
	" berry wines ",
	" besan flour ",
	" big carrots ",
	" bing cherry ",
	" bird chiles ",
	" bird pepper ",
	" biscuit mix ",
	" black beans ",
	" blackberrys ",
	" blackcurrant ",
	" blended oil ",
	" bleu cheese ",
	" blood orange ",
	" blue cheese ",
	" blueberries ",
	" bocconcinis ",
	" bomba rices ",
	" bourbon rye ",
	" boysenberry ",
	" boysenberrys ",
	" brazil nuts ",
	" bread crumb ",
	" bread flour ",
	" bread rolls ",
	" breadcrumbs ",
	" breadcrumbs ",
	" breadfruits ",
	" brie cheese ",
	" brioche bun ",
	" broad beans ",
	" broad beanss ",
	" broccoflower ",
	" broccolinis ",
	" brook trout ",
	" brown onion ",
	" brown rices ",
	" brown sugar ",
	" bruschettum ",
	" bulb fennel ",
	" bulb garlic ",
	" burger buns ",
	" butter bean ",
	" butter oils ",
	" buttermilks ",
	" cake flours ",
	" cake yeasts ",
	" calf livers ",
	" californium ",
	" campanelles ",
	" canary melon ",
	" candy canes ",
	" cane sugars ",
	" cane syrups ",
	" cannellinis ",
	" cannellonis ",
	" canola oils ",
	" cantaloupes ",
	" caper berry ",
	" caper brine ",
	" caper juice ",
	" caperberrys ",
	" cappuccinos ",
	" carne asada ",
	" carob chips ",
	" carrot curl ",
	" cashew half ",
	" cashew milk ",
	" cashew nuts ",
	" cauliflower ",
	" cauliflowers ",
	" celery ribs ",
	" celery salt ",
	" celery seed ",
	" celery seeds ",
	" channa dals ",
	" chanterelle ",
	" chardonnays ",
	" chat masala ",
	" cheese curd ",
	" cheese soup ",
	" cheese whiz ",
	" cheesecakes ",
	" cheesecloth ",
	" cherry cola ",
	" cherry jams ",
	" chex cereal ",
	" chicken fat ",
	" chile salsa ",
	" chile sauce ",
	" chili beans ",
	" chili pepper ",
	" chili sauce ",
	" chimichurri ",
	" chium seeds ",
	" chive dills ",
	" chuck roast ",
	" chuck steak ",
	" citric acid ",
	" citru fruit ",
	" citru juice ",
	" citru vodka ",
	" clam broths ",
	" clam juices ",
	" clear honey ",
	" clementines ",
	" co lettuces ",
	" coconut oil ",
	" coconut rum ",
	" coffee bean ",
	" coffee cake ",
	" coho salmon ",
	" comice pear ",
	" comt cheese ",
	" cooking oil ",
	" cool waters ",
	" corn breads ",
	" corn flours ",
	" corn kernel ",
	" corn niblet ",
	" corn relish ",
	" corn starch ",
	" corn syrups ",
	" corned beef ",
	" cornish hen ",
	" cornstarchs ",
	" corona beer ",
	" cottonseeds ",
	" country ham ",
	" cranberries ",
	" cream sauce ",
	" crisp apple ",
	" crisp bacon ",
	" crme frache ",
	" crusty roll ",
	" cumin seeds ",
	" cured meats ",
	" curing salt ",
	" curly kales ",
	" custard mix ",
	" cut chicken ",
	" cuttlefishs ",
	" date sugars ",
	" date syrups ",
	" decorations ",
	" deli turkey ",
	" demi glaces ",
	" dill chives ",
	" dill fronds ",
	" dill pickle ",
	" dilly beans ",
	" dinner roll ",
	" dragonfruits ",
	" dry mustard ",
	" dry sherrys ",
	" duck confit ",
	" edam cheese ",
	" egg noodles ",
	" egg noodles ",
	" egg tomatos ",
	" egg whisked ",
	" elderberrys ",
	" emmentalers ",
	" english pea ",
	" erythritols ",
	" farfallines ",
	" fennel bulb ",
	" fennel seed ",
	" feta cheese ",
	" fettuccines ",
	" filo pastry ",
	" fino sherry ",
	" fish broths ",
	" fish sauces ",
	" fish steaks ",
	" flaky salts ",
	" flank steak ",
	" flat breads ",
	" frangelicos ",
	" frankfurter ",
	" french bean ",
	" french frys ",
	" french roll ",
	" fried onion ",
	" fried rices ",
	" fried tofus ",
	" fruit juice ",
	" fruit salad ",
	" fruit salsa ",
	" frying oils ",
	" gala apples ",
	" garden peas ",
	" garlic bulb ",
	" garlic oils ",
	" garlic salt ",
	" gem lettuce ",
	" ginger ales ",
	" ginger beer ",
	" ginger snap ",
	" gingerbread ",
	" gingerroots ",
	" gingersnaps ",
	" gluten free ",
	" goat cheese ",
	" goji berrys ",
	" golden beet ",
	" golden rums ",
	" good breads ",
	" good honeys ",
	" gooseberrys ",
	" gorgonzolas ",
	" granola bar ",
	" grape jelly ",
	" grapefruits ",
	" greek fetum ",
	" greek olive ",
	" green apple ",
	" green beans ",
	" green beans ",
	" green chile ",
	" green chili ",
	" green grape ",
	" green mango ",
	" green olive ",
	" green onion ",
	" green pears ",
	" green pepper ",
	" green salsa ",
	" guava juice ",
	" gumbo files ",
	" haa avocado ",
	" hardy green ",
	" hash browns ",
	" heavy cream ",
	" herb butter ",
	" hoagie buns ",
	" hollandaise ",
	" holy basils ",
	" hominy grit ",
	" honey syrup ",
	" horseradish ",
	" horseradishs ",
	" huckleberry ",
	" huckleberrys ",
	" iced waters ",
	" icing sugar ",
	" inch celery ",
	" inch ginger ",
	" irish cream ",
	" jack cheese ",
	" jamaica rum ",
	" jelly beans ",
	" kaffir lime ",
	" kaiser roll ",
	" kecap manis ",
	" kernel corn ",
	" ketjap mani ",
	" kidney bean ",
	" kidney beans ",
	" king salmon ",
	" kiwi fruits ",
	" kosher salt ",
	" lady apples ",
	" ladyfingers ",
	" lager beers ",
	" lamb minces ",
	" lamb steaks ",
	" langoustine ",
	" leafy green ",
	" leek greens ",
	" leek whites ",
	" lemon aioli ",
	" lemon balms ",
	" lemon basil ",
	" lemon curds ",
	" lemon grass ",
	" lemon grasss ",
	" lemon halfs ",
	" lemon juice ",
	" lemon limes ",
	" lemon myrtle ",
	" lemon peels ",
	" lemon rinds ",
	" lemon soles ",
	" lemon thyme ",
	" lemon vodka ",
	" lemon wheel ",
	" lemongrasss ",
	" lime butter ",
	" lime halved ",
	" lime juices ",
	" lime wheels ",
	" limoncellos ",
	" lingonberry ",
	" live mussel ",
	" liver pates ",
	" liverwursts ",
	" loaf breads ",
	" long peppers ",
	" lowfat milk ",
	" madra curry ",
	" maldon salt ",
	" malibu rums ",
	" malt syrups ",
	" mango pulps ",
	" mango puree ",
	" mango salsa ",
	" mango-ginger ",
	" manila clam ",
	" maple sugar ",
	" maple syrup ",
	" mara pepper ",
	" margaritums ",
	" marrow bone ",
	" marshmallow ",
	" masa harina ",
	" mascarpones ",
	" matzo meals ",
	" matzoh meal ",
	" mayonnaises ",
	" meat sauces ",
	" meyer lemon ",
	" microgreens ",
	" milk creams ",
	" milk crumbs ",
	" milk yogurt ",
	" minestrones ",
	" mint leaves ",
	" mint pestos ",
	" mint sauces ",
	" mint syrups ",
	" minute rice ",
	" mixed berry ",
	" mixed citru ",
	" mixed fruit ",
	" mixed green ",
	" mixed herbs ",
	" mixed olive ",
	" mixed seeds ",
	" mixed spice ",
	" mole sauces ",
	" mortadellas ",
	" mozzarellas ",
	" muscat wine ",
	" mustard oil ",
	" naan breads ",
	" neutral oil ",
	" new potatos ",
	" nilla wafer ",
	" noilly prat ",
	" nonfat milk ",
	" nori sheets ",
	" nut biscuit ",
	" nut butters ",
	" oil butters ",
	" onion salts ",
	" opal basils ",
	" orange oils ",
	" orange peel ",
	" orange soda ",
	" orecchiette ",
	" oreo cookie ",
	" oz parmesan ",
	" oz prosecco ",
	" pacific cod ",
	" paella rice ",
	" pale lagers ",
	" palm sugars ",
	" pancake mix ",
	" papaya seed ",
	" pappardelle ",
	" parmigianos ",
	" passionfruit ",
	" pat butters ",
	" pea tendril ",
	" peach puree ",
	" peach syrup ",
	" peanut oils ",
	" pear brandy ",
	" pear juices ",
	" pear nectar ",
	" pear tomato ",
	" pearl onion ",
	" pearl sugar ",
	" pecan meals ",
	" pekin ducks ",
	" peking duck ",
	" peperoncini ",
	" peperoncino ",
	" peppercorns ",
	" peppermints ",
	" perciatelli ",
	" pesto sauce ",
	" petite peas ",
	" pickle chip ",
	" pie pastrys ",
	" pig trotter ",
	" pigeon peas ",
	" pignoli nut ",
	" piloncillos ",
	" pinot noirs ",
	" pinto beans ",
	" pitum bread ",
	" pizza crust ",
	" pizza dough ",
	" pizza sauce ",
	" plum sauces ",
	" plum tomato ",
	" poached egg ",
	" pomegranate ",
	" pomegranates ",
	" ponzu sauce ",
	" popped corn ",
	" poppy seeds ",
	" pork bellys ",
	" pork livers ",
	" pork roasts ",
	" pork steaks ",
	" portobellos ",
	" potato buns ",
	" potato chip ",
	" potato roll ",
	" prosciuttos ",
	" prune juice ",
	" prune plums ",
	" prune puree ",
	" pudding mix ",
	" puff pastry ",
	" puffed rice ",
	" pumpkin pie ",
	" pumpkinseed ",
	" purple plum ",
	" quaker oats ",
	" quatre pice ",
	" quesadillas ",
	" quick bread ",
	" ramp greens ",
	" raspberries ",
	" real butter ",
	" red peppers ",
	" red peppers ",
	" rib celerys ",
	" rice cereal ",
	" rice flours ",
	" rice krispy ",
	" rice noodle ",
	" rice papers ",
	" rice pilafs ",
	" rittenhouse ",
	" roast beefs ",
	" roast porks ",
	" rolled oats ",
	" roma tomato ",
	" romano bean ",
	" rome apples ",
	" rose petals ",
	" rose waters ",
	" royal icing ",
	" rubbed sage ",
	" rum liqueur ",
	" rump roasts ",
	" rump steaks ",
	" runner beans ",
	" rye whiskey ",
	" rye whiskys ",
	" sage thymes ",
	" salad green ",
	" salal berrys ",
	" salmon roes ",
	" salmon skin ",
	" salsa verde ",
	" salt pepper ",
	" salt waters ",
	" san marzano ",
	" sarsaparilla ",
	" satay sauce ",
	" sauerkrauts ",
	" season salt ",
	" seckel pear ",
	" serrano ham ",
	" sesame oils ",
	" sesame seed ",
	" shallot oil ",
	" sheet matzo ",
	" sherry wine ",
	" shiro misos ",
	" shortbreads ",
	" shortenings ",
	" silken tofu ",
	" sirloin tip ",
	" skate wings ",
	" skirt steak ",
	" slider buns ",
	" smith apple ",
	" smoked fish ",
	" smoked hams ",
	" smoked salt ",
	" smoked tofu ",
	" smoky bacon ",
	" snack cakes ",
	" soba noodle ",
	" soda breads ",
	" soda waters ",
	" soft butter ",
	" soft cheese ",
	" soft drinks ",
	" soft fruits ",
	" solid tunas ",
	" sour apples ",
	" sour cherry ",
	" sour creams ",
	" sour pickle ",
	" soy protein ",
	" soy yogurts ",
	" soya sauces ",
	" soybean oil ",
	" spaghettini ",
	" spelt berry ",
	" spelt flour ",
	" spiced rums ",
	" spicy green ",
	" sponge cake ",
	" squash seed ",
	" st germains ",
	" stale bread ",
	" star anises ",
	" star fruits ",
	" steak sauce ",
	" sticky rice ",
	" stout beers ",
	" strawberrys ",
	" string bean ",
	" strip bacon ",
	" strip steak ",
	" stroganoffs ",
	" sugar canes ",
	" sugar cones ",
	" sugar syrup ",
	" sushi rices ",
	" sweet onion ",
	" sweet potato ",
	" sweetbreads ",
	" swiss chard ",
	" table cream ",
	" table salts ",
	" tagliatelle ",
	" tart apples ",
	" tart cherry ",
	" tart crusts ",
	" tawny ports ",
	" teff flours ",
	" tender herb ",
	" tepid water ",
	" thai basils ",
	" thai chiles ",
	" thai chilis ",
	" thai pepper ",
	" thick bacon ",
	" thyme honey ",
	" tiger prawn ",
	" tiny capers ",
	" tiny potato ",
	" tium marium ",
	" toast point ",
	" toffee bits ",
	" tomato jams ",
	" tomato oils ",
	" tomato soup ",
	" tonic water ",
	" torn basils ",
	" torn greens ",
	" tortellinis ",
	" tortiglioni ",
	" triple secs ",
	" truffle oil ",
	" tuna salads ",
	" tuna steaks ",
	" tuna waters ",
	" turkey bone ",
	" turkey fats ",
	" turkey hams ",
	" turkey meat ",
	" turkey wing ",
	" udon noodle ",
	" vanilla pod ",
	" veal roasts ",
	" vegan sugar ",
	" vermicellis ",
	" vinaigrette ",
	" vsop cognac ",
	" walnut half ",
	" walnut oils ",
	" watercresss ",
	" watermelons ",
	" waxy potato ",
	" weisswursts ",
	" wheat bagel ",
	" wheat beers ",
	" wheat berry ",
	" wheat brans ",
	" wheat bread ",
	" wheat flour ",
	" wheat germs ",
	" wheat penne ",
	" wheat pitas ",
	" wheat pitum ",
	" wheat toast ",
	" white beans ",
	" white beers ",
	" white corns ",
	" white fishs ",
	" white flour ",
	" white grape ",
	" white grits ",
	" white karos ",
	" white misos ",
	" white onion ",
	" white peach ",
	" white ports ",
	" white radish ",
	" white rices ",
	" white sauce ",
	" white sugar ",
	" white tunas ",
	" white wheat ",
	" white wines ",
	" wild salmon ",
	" wonton skin ",
	" xanthan gum ",
	" yellow beet ",
	" yellow corn ",
	" yellow finn ",
	" yellow grit ",
	" yellow miso ",
	" yellow rice ",
	" yukon golds ",
	" yuzu juices ",
	" ziti pennes ",
	" agar agars ",
	" almond oil ",
	" amber beer ",
	" amber rums ",
	" andouilles ",
	" anise oils ",
	" anise seed ",
	" anjou pear ",
	" apple cake ",
	" apple jams ",
	" apple pear ",
	" apple pies ",
	" applejacks ",
	" applesauce ",
	" arrowroots ",
	" artichokes ",
	" asafetidas ",
	" asafoetida ",
	" asafoetidas ",
	" asian pear ",
	" asparaguss ",
	" aubergines ",
	" azuki beans ",
	" baby beets ",
	" baby caper ",
	" baby clams ",
	" baby corns ",
	" baby green ",
	" baby kales ",
	" baby lambs ",
	" baby leeks ",
	" baby onion ",
	" bacon bits ",
	" bacon fats ",
	" bagel chip ",
	" baked bean ",
	" baked hams ",
	" baked tofu ",
	" baking mix ",
	" banana rum ",
	" basil mint ",
	" basil oils ",
	" bay leaves ",
	" bay shrimp ",
	" bbq sauces ",
	" bean broth ",
	" bean puree ",
	" bean salad ",
	" bean sprout ",
	" beaten egg ",
	" beef bones ",
	" beef broth ",
	" beef cheek ",
	" beef chuck ",
	" beef liver ",
	" beef mince ",
	" beef roast ",
	" beef steak ",
	" beef stews ",
	" beef stock ",
	" beef suets ",
	" beefsteaks ",
	" beet green ",
	" bell pepper ",
	" benne seed ",
	" berry wine ",
	" big carrot ",
	" bird chile ",
	" black beans ",
	" blackberry ",
	" blackberrys ",
	" blackfishs ",
	" blue crabs ",
	" blueberrys ",
	" bocconcini ",
	" bologneses ",
	" bomba rice ",
	" bosc pears ",
	" boysenberry ",
	" bratwursts ",
	" brazil nut ",
	" bread roll ",
	" breadcrumb ",
	" breadfruit ",
	" breadfruits ",
	" breakfasts ",
	" broad bean ",
	" broad beans ",
	" broccolini ",
	" brown rice ",
	" buckwheats ",
	" budweisers ",
	" burger bun ",
	" butter oil ",
	" buttermilk ",
	" butternuts ",
	" cacao nibs ",
	" cake flour ",
	" cake tofus ",
	" cake yeast ",
	" calamatums ",
	" calf liver ",
	" camemberts ",
	" campanelle ",
	" candlenuts ",
	" candy bars ",
	" candy cane ",
	" cane sugar ",
	" cane syrup ",
	" cannellini ",
	" cannelloni ",
	" canola oil ",
	" cantaloupe ",
	" cantaloupes ",
	" capellinis ",
	" caperberry ",
	" capocollos ",
	" caponatums ",
	" cappuccino ",
	" carambolas ",
	" cardamaros ",
	" caribbeans ",
	" carob chip ",
	" cashew nut ",
	" cauliflower ",
	" cavatappis ",
	" cavatellis ",
	" celery rib ",
	" celery seed ",
	" chamomiles ",
	" champagnes ",
	" channa dal ",
	" chardonnay ",
	" cheesecake ",
	" chermoulas ",
	" cherry jam ",
	" chick peas ",
	" chickpeass ",
	" chile oils ",
	" chili bean ",
	" chili oils ",
	" chium seed ",
	" chive dill ",
	" chive plus ",
	" chocolates ",
	" chopsticks ",
	" chow meins ",
	" ciabattums ",
	" clam broth ",
	" clam juice ",
	" clementine ",
	" clementines ",
	" cloudberrys ",
	" club sodas ",
	" co lettuce ",
	" coca colas ",
	" cocoa nibs ",
	" cod steaks ",
	" cointreaus ",
	" condiments ",
	" cool water ",
	" cool whips ",
	" corianders ",
	" corn bread ",
	" corn chips ",
	" corn flour ",
	" corn grits ",
	" corn husks ",
	" corn meals ",
	" corn salads ",
	" corn syrup ",
	" cornbreads ",
	" cornflakes ",
	" cornflours ",
	" cornichons ",
	" cornstarch ",
	" cottonseed ",
	" courgettes ",
	" crab cakes ",
	" crab meats ",
	" cranberrys ",
	" croissants ",
	" cumin seed ",
	" cured hams ",
	" cured meat ",
	" curly kale ",
	" curry leafs ",
	" curry oils ",
	" cuttlefish ",
	" dandelions ",
	" date sugar ",
	" date syrup ",
	" decoration ",
	" deli meats ",
	" demi glace ",
	" dill chive ",
	" dill frond ",
	" dill seeds ",
	" dill weeds ",
	" dilly bean ",
	" dr peppers ",
	" dragonfruit ",
	" dried limes ",
	" dry sherry ",
	" egg noodle ",
	" egg salads ",
	" egg tomato ",
	" egg whites ",
	" elderberry ",
	" elderberrys ",
	" emmentaler ",
	" emmenthals ",
	" enchiladas ",
	" entrecotes ",
	" erythritol ",
	" escabeches ",
	" everclears ",
	" farfalline ",
	" fava beans ",
	" fenugreeks ",
	" fettuccine ",
	" fettucines ",
	" fiddleheads ",
	" field peas ",
	" fingerroots ",
	" fish bones ",
	" fish broth ",
	" fish cakes ",
	" fish sauce ",
	" fish steak ",
	" flaky salt ",
	" flat bread ",
	" flatbreads ",
	" flavorings ",
	" flax meals ",
	" flax seeds ",
	" focacciums ",
	" framboises ",
	" frangelico ",
	" french fry ",
	" fried eggs ",
	" fried rice ",
	" fried tofu ",
	" frittatums ",
	" fruit jams ",
	" frying oil ",
	" gala apple ",
	" garden pea ",
	" garlic oil ",
	" gem squashs ",
	" ginger ale ",
	" gingerroot ",
	" gingersnap ",
	" goat meats ",
	" goat milks ",
	" gochugarus ",
	" gochujangs ",
	" goji berry ",
	" goji berrys ",
	" golden rum ",
	" good bread ",
	" good honey ",
	" goose fats ",
	" gooseberry ",
	" gooseberrys ",
	" gorgonzola ",
	" grapefruit ",
	" grapefruits ",
	" grapeseeds ",
	" green bean ",
	" green beans ",
	" green peas ",
	" green pear ",
	" grenadines ",
	" guacamoles ",
	" guanciales ",
	" gumbo file ",
	" half halfs ",
	" ham steaks ",
	" hamburgers ",
	" hash brown ",
	" hazelnutss ",
	" heath bars ",
	" hemp seeds ",
	" herb salts ",
	" herbsaints ",
	" hero rolls ",
	" hoagie bun ",
	" hoja santas ",
	" holy basil ",
	" holy basils ",
	" honeycombs ",
	" horseradish ",
	" hot sauces ",
	" huckleberry ",
	" iced water ",
	" jackfruits ",
	" jambalayas ",
	" jelly bean ",
	" jumbo eggs ",
	" kecap mani ",
	" kidney bean ",
	" kiwi fruit ",
	" kiwi fruits ",
	" kiwifruits ",
	" kochujangs ",
	" kochukarus ",
	" lady apple ",
	" ladyfinger ",
	" lager beer ",
	" lamb bones ",
	" lamb meats ",
	" lamb mince ",
	" lamb steak ",
	" lamb stews ",
	" leek green ",
	" leek white ",
	" lemon balm ",
	" lemon balms ",
	" lemon curd ",
	" lemon grass ",
	" lemon half ",
	" lemon lime ",
	" lemon peel ",
	" lemon rind ",
	" lemon sole ",
	" lemongrass ",
	" lima beans ",
	" lime halfs ",
	" lime juice ",
	" lime peels ",
	" lime wheel ",
	" limoncello ",
	" liver pate ",
	" liverwurst ",
	" loaf bread ",
	" long beans ",
	" long pepper ",
	" mahi mahis ",
	" malibu rum ",
	" malt syrup ",
	" mango pulp ",
	" mango rums ",
	" manicottis ",
	" margarines ",
	" margaritum ",
	" marmalades ",
	" mascarpone ",
	" matzo meal ",
	" mayonnaise ",
	" meat loafs ",
	" meat sauce ",
	" medallions ",
	" membrillos ",
	" microgreen ",
	" milk cream ",
	" milk crumb ",
	" minestrone ",
	" mint pesto ",
	" mint sauce ",
	" mint syrup ",
	" mixed herb ",
	" mixed nuts ",
	" mixed seed ",
	" mole sauce ",
	" mortadella ",
	" mozzarella ",
	" mung beans ",
	" mushroomss ",
	" naan bread ",
	" navy beans ",
	" nectarines ",
	" new potato ",
	" newspapers ",
	" nori sheet ",
	" nuoc chams ",
	" nut butter ",
	" oat flours ",
	" oil butter ",
	" oil sprays ",
	" olive oils ",
	" onion salt ",
	" opal basil ",
	" orange oil ",
	" pale lager ",
	" palm sugar ",
	" pan juices ",
	" pan sprays ",
	" pancettums ",
	" panettones ",
	" parchments ",
	" parma hams ",
	" parmigiano ",
	" partridges ",
	" pat butter ",
	" peach jams ",
	" peanut oil ",
	" pear juice ",
	" pecan meal ",
	" pecan nuts ",
	" pekin duck ",
	" peppercorn ",
	" peppermint ",
	" pepperonis ",
	" persimmons ",
	" petite pea ",
	" pie crusts ",
	" pie doughs ",
	" pie pastry ",
	" pie shells ",
	" pigeon pea ",
	" piloncillo ",
	" pine nutss ",
	" pineapples ",
	" pinot noir ",
	" pinto bean ",
	" pinto beans ",
	" piri piris ",
	" pistachios ",
	" plum sauce ",
	" plum wines ",
	" pole beans ",
	" pomegranate ",
	" poppy seed ",
	" poppy seeds ",
	" pork belly ",
	" pork butts ",
	" pork lards ",
	" pork liver ",
	" pork meats ",
	" pork roast ",
	" pork steak ",
	" port wines ",
	" portobello ",
	" pot roasts ",
	" potato bun ",
	" prosciutto ",
	" provolones ",
	" prune plum ",
	" quail eggs ",
	" quaker oat ",
	" quesadilla ",
	" quick oats ",
	" radiatores ",
	" radicchios ",
	" ramp green ",
	" raspberrys ",
	" reblochons ",
	" red onions ",
	" red pepper ",
	" redcurrants ",
	" remoulades ",
	" rib celery ",
	" rib roasts ",
	" rib steaks ",
	" rice cakes ",
	" rice flour ",
	" rice milks ",
	" rice paper ",
	" rice pilaf ",
	" rice wines ",
	" roast beef ",
	" roast pork ",
	" rock melons ",
	" rock salts ",
	" rolled oat ",
	" rome apple ",
	" roqueforts ",
	" rose petal ",
	" rose water ",
	" rose wines ",
	" rosewaters ",
	" ruby ports ",
	" rump roast ",
	" rump steak ",
	" runner bean ",
	" rye berrys ",
	" rye breads ",
	" rye flours ",
	" rye whisky ",
	" sablefishs ",
	" safflowers ",
	" sage thyme ",
	" salad oils ",
	" salal berry ",
	" salmon roe ",
	" salt porks ",
	" salt water ",
	" sauerkraut ",
	" savoiardis ",
	" scallionss ",
	" schnitzels ",
	" sesame oil ",
	" shellfishs ",
	" shiro miso ",
	" short ribs ",
	" shortbread ",
	" shortcakes ",
	" shortening ",
	" skate wing ",
	" skim milks ",
	" slider bun ",
	" sloppy jos ",
	" smoked ham ",
	" snack cake ",
	" snap beans ",
	" soda bread ",
	" soda water ",
	" soft drink ",
	" soft fruit ",
	" soft herbs ",
	" soft rolls ",
	" soft tofus ",
	" solid tuna ",
	" sour apple ",
	" sour cream ",
	" sourdoughs ",
	" soy creams ",
	" soy flours ",
	" soy sauces ",
	" soy yogurt ",
	" soya milks ",
	" soya sauce ",
	" spaghettis ",
	" spare ribs ",
	" spearmints ",
	" spice mixs ",
	" spice rubs ",
	" spiced rum ",
	" spirulinas ",
	" split peas ",
	" squid inks ",
	" st germain ",
	" star anise ",
	" star anises ",
	" star fruit ",
	" star fruits ",
	" stew beefs ",
	" stew meats ",
	" stout beer ",
	" strawberry ",
	" strawberrys ",
	" stroganoff ",
	" succotashs ",
	" sugar cane ",
	" sugar cone ",
	" sunflowers ",
	" sushi rice ",
	" sweetbread ",
	" sweetcorns ",
	" swordfishs ",
	" table salt ",
	" tamarillos ",
	" tangerines ",
	" tap waters ",
	" tart apple ",
	" tart crust ",
	" tater tots ",
	" tawny port ",
	" teff flour ",
	" thai basil ",
	" thai basils ",
	" thai chile ",
	" thai chili ",
	" thai crabs ",
	" tiny caper ",
	" toffee bit ",
	" togarashis ",
	" tomatillos ",
	" tomato jam ",
	" tomato oil ",
	" tonka beans ",
	" topinamburs ",
	" torn basil ",
	" torn green ",
	" tortellini ",
	" trail mixs ",
	" triple sec ",
	" trout roes ",
	" tuna fishs ",
	" tuna salad ",
	" tuna steak ",
	" tuna water ",
	" turbinados ",
	" turkey fat ",
	" turkey ham ",
	" ugli fruits ",
	" veal bones ",
	" veal roast ",
	" vegenaises ",
	" vegetables ",
	" vermicelli ",
	" vin santos ",
	" walnut oil ",
	" watercress ",
	" watercresss ",
	" watermelon ",
	" watermelons ",
	" weisswurst ",
	" wheat beer ",
	" wheat bran ",
	" wheat buns ",
	" wheat germ ",
	" wheat pita ",
	" white bean ",
	" white beer ",
	" white corn ",
	" white fish ",
	" white grit ",
	" white karo ",
	" white miso ",
	" white port ",
	" white rice ",
	" white rums ",
	" white tuna ",
	" white wine ",
	" whitefishs ",
	" wild boars ",
	" wild ducks ",
	" wild rices ",
	" wood chips ",
	" yukon gold ",
	" yuzu juice ",
	" zinfandels ",
	" ziti penne ",
	" absinthes ",
	" advocaats ",
	" agar agar ",
	" aged rums ",
	" ahi tunas ",
	" allspices ",
	" amaranths ",
	" amarettis ",
	" amarettos ",
	" amber rum ",
	" andouille ",
	" anise oil ",
	" anisettes ",
	" apple jam ",
	" apple pie ",
	" applejack ",
	" armagnacs ",
	" arrowroot ",
	" artichoke ",
	" artichokes ",
	" asafetida ",
	" asafoetida ",
	" asparagus ",
	" asparaguss ",
	" assemblys ",
	" aubergine ",
	" aubergines ",
	" azuki bean ",
	" baby beet ",
	" baby clam ",
	" baby corn ",
	" baby kale ",
	" baby lamb ",
	" baby leek ",
	" baby peas ",
	" bacon bit ",
	" bacon fat ",
	" baguettes ",
	" baked ham ",
	" balsamics ",
	" barbecues ",
	" barberrys ",
	" basil oil ",
	" bay leafs ",
	" bbq sauce ",
	" bechamels ",
	" beef bone ",
	" beef fats ",
	" beef ribs ",
	" beef stew ",
	" beef suet ",
	" beefsteak ",
	" beetroots ",
	" biscottis ",
	" bisquicks ",
	" black bean ",
	" blackberry ",
	" blackfish ",
	" blue crab ",
	" blueberry ",
	" blueberrys ",
	" bluefishs ",
	" bok choys ",
	" bolognese ",
	" bosc pear ",
	" bouillons ",
	" branzinos ",
	" bratwurst ",
	" breadfruit ",
	" breakfast ",
	" bresaolas ",
	" broccolis ",
	" bucatinis ",
	" buckwheat ",
	" budweiser ",
	" burratums ",
	" butternut ",
	" cacao nib ",
	" cake mixs ",
	" cake tofu ",
	" calabreses ",
	" calamaris ",
	" calamatum ",
	" callaloos ",
	" camembert ",
	" candlenut ",
	" candy bar ",
	" cantaloupe ",
	" capellini ",
	" capicolas ",
	" capocollo ",
	" caponatum ",
	" capsicums ",
	" carambola ",
	" cardamaro ",
	" cardamoms ",
	" caribbean ",
	" cavatappi ",
	" cavatelli ",
	" celeriacs ",
	" chambords ",
	" chamomile ",
	" chamomiles ",
	" champagne ",
	" char sius ",
	" cherimoyas ",
	" chermoula ",
	" cherriess ",
	" chestnuts ",
	" chick pea ",
	" chickpeas ",
	" chickpeas ",
	" chile oil ",
	" chili oil ",
	" chipotles ",
	" chive plu ",
	" chocolate ",
	" chopstick ",
	" chow mein ",
	" choy sums ",
	" ciabattum ",
	" cilantros ",
	" cinnamons ",
	" clementine ",
	" cloudberry ",
	" club soda ",
	" coca cola ",
	" cocktails ",
	" cocoa nib ",
	" cod steak ",
	" cointreau ",
	" coleslaws ",
	" condiment ",
	" cool whip ",
	" coriander ",
	" corn chip ",
	" corn cobs ",
	" corn grit ",
	" corn husk ",
	" corn meal ",
	" corn oils ",
	" corn salad ",
	" cornbread ",
	" cornflake ",
	" cornflour ",
	" cornichon ",
	" cornmeals ",
	" courgette ",
	" courgettes ",
	" crab cake ",
	" crab meat ",
	" crabmeats ",
	" cranberry ",
	" cranberrys ",
	" crawfishs ",
	" croissant ",
	" crostinis ",
	" cucumbers ",
	" culantros ",
	" cured ham ",
	" curry leaf ",
	" curry oil ",
	" dandelion ",
	" deli hams ",
	" deli meat ",
	" demeraras ",
	" dill seed ",
	" dill seeds ",
	" dill weed ",
	" ditalinis ",
	" dr pepper ",
	" drambuies ",
	" dressings ",
	" dried lime ",
	" dubonnets ",
	" duck eggs ",
	" duck fats ",
	" ducklings ",
	" dumplings ",
	" ear corns ",
	" egg rolls ",
	" egg salad ",
	" egg washs ",
	" egg white ",
	" egg yolks ",
	" eggplants ",
	" elderberry ",
	" emmentals ",
	" emmenthal ",
	" empanadas ",
	" enchilada ",
	" entrecote ",
	" escabeche ",
	" escaroles ",
	" espressos ",
	" everclear ",
	" falernums ",
	" farfalles ",
	" fava bean ",
	" fenugreek ",
	" fenugreeks ",
	" fettucine ",
	" fiddlehead ",
	" field pea ",
	" fingerroot ",
	" fish bone ",
	" fish cake ",
	" fish oils ",
	" flatbread ",
	" flavoring ",
	" flax meal ",
	" flax seed ",
	" flaxseeds ",
	" flounders ",
	" focaccium ",
	" foie gras ",
	" framboise ",
	" fried egg ",
	" frittatum ",
	" frostings ",
	" fructoses ",
	" fruit jam ",
	" furikakes ",
	" galangals ",
	" gelatines ",
	" gem squash ",
	" giandujas ",
	" goat meat ",
	" goat milk ",
	" gochugaru ",
	" gochujang ",
	" goji berry ",
	" gold rums ",
	" goose fat ",
	" gooseberry ",
	" grapefruit ",
	" grapeseed ",
	" green bean ",
	" green pea ",
	" grenadine ",
	" guacamole ",
	" guajillos ",
	" guanciale ",
	" guar gums ",
	" guinnesss ",
	" habaneros ",
	" half half ",
	" halloumis ",
	" ham bones ",
	" ham hocks ",
	" ham steak ",
	" hamburger ",
	" hazelnuts ",
	" hazelnuts ",
	" heath bar ",
	" hemp seed ",
	" herb salt ",
	" herbsaint ",
	" hero roll ",
	" hoja santa ",
	" holy basil ",
	" honeycomb ",
	" honeydews ",
	" hot sauce ",
	" irm tofus ",
	" jackfruit ",
	" jackfruits ",
	" jalapenos ",
	" jambalaya ",
	" jumbo egg ",
	" key limes ",
	" kielbasas ",
	" kinh giois ",
	" kiwi fruit ",
	" kiwifruit ",
	" kochujang ",
	" kochukaru ",
	" kohlrabis ",
	" kool aids ",
	" lamb bone ",
	" lamb meat ",
	" lamb ribs ",
	" lamb stew ",
	" lavenders ",
	" lecithins ",
	" lemon balm ",
	" lemon zes ",
	" lemonades ",
	" lima bean ",
	" lima beans ",
	" lime half ",
	" lime peel ",
	" linguines ",
	" linguinis ",
	" liquorices ",
	" lollipops ",
	" long bean ",
	" macaronis ",
	" macaroons ",
	" mackerels ",
	" mahi mahi ",
	" mahimahis ",
	" manchegos ",
	" mandarins ",
	" mandarines ",
	" mangetouts ",
	" mango rum ",
	" manicotti ",
	" margarine ",
	" marinades ",
	" marjorams ",
	" marmalade ",
	" marzipans ",
	" meat cuts ",
	" meat loaf ",
	" meatballs ",
	" meatloafs ",
	" medallion ",
	" membrillo ",
	" meringues ",
	" mirepoixs ",
	" mirlitons ",
	" mixed nut ",
	" molassess ",
	" monkfishs ",
	" mulberrys ",
	" mung bean ",
	" mung beans ",
	" mung dals ",
	" mushrooms ",
	" mushrooms ",
	" navy bean ",
	" navy beans ",
	" nectarine ",
	" nectarines ",
	" newspaper ",
	" nopalitos ",
	" nuoc cham ",
	" nuoc nams ",
	" nut milks ",
	" oat brans ",
	" oat flour ",
	" oil spray ",
	" okra pods ",
	" olive oil ",
	" oz bacons ",
	" oz ndujas ",
	" paccheris ",
	" pak chois ",
	" palm oils ",
	" pan juice ",
	" pan spray ",
	" pancettum ",
	" panettone ",
	" parchment ",
	" parma ham ",
	" parmesans ",
	" partridge ",
	" passatums ",
	" pastramis ",
	" patty pans ",
	" peach jam ",
	" pecan nut ",
	" pecorinos ",
	" pepperoni ",
	" persimmon ",
	" persimmons ",
	" pheasants ",
	" pie crust ",
	" pie dough ",
	" pie shell ",
	" pig tails ",
	" piment ns ",
	" pimentons ",
	" pimientos ",
	" pine nuts ",
	" pine nuts ",
	" pineapple ",
	" pineapples ",
	" pinto bean ",
	" piri piri ",
	" pistachio ",
	" plantains ",
	" plum jams ",
	" plum wine ",
	" pole bean ",
	" polentums ",
	" poppy seed ",
	" popsicles ",
	" pork butt ",
	" pork fats ",
	" pork lard ",
	" pork meat ",
	" pork ribs ",
	" port wine ",
	" pot roast ",
	" potatoess ",
	" preserves ",
	" proseccos ",
	" provolone ",
	" purslanes ",
	" quail egg ",
	" quick oat ",
	" raclettes ",
	" radiatore ",
	" radicchio ",
	" radicchios ",
	" raspberry ",
	" raspberrys ",
	" reblochon ",
	" red onion ",
	" red wines ",
	" redcurrant ",
	" remoulade ",
	" rib roast ",
	" rib steak ",
	" rice cake ",
	" rice milk ",
	" rice mixs ",
	" rice wine ",
	" rieslings ",
	" rigatonis ",
	" rock melon ",
	" rock salt ",
	" roquefort ",
	" rose wine ",
	" rosemarys ",
	" rosewater ",
	" ruby port ",
	" rutabagas ",
	" rye berry ",
	" rye bread ",
	" rye flour ",
	" sablefish ",
	" safflower ",
	" salad oil ",
	" salt cods ",
	" salt plus ",
	" salt pork ",
	" sandwichs ",
	" sangriums ",
	" sassafras ",
	" sassafrass ",
	" sauternes ",
	" savoiardi ",
	" scallions ",
	" scallions ",
	" schmaltzs ",
	" schnitzel ",
	" semolinas ",
	" shallotss ",
	" shaoxings ",
	" shellfish ",
	" shiitakes ",
	" short rib ",
	" shortcake ",
	" skim milk ",
	" sloppy jo ",
	" smoothies ",
	" snap bean ",
	" snap peas ",
	" snow peas ",
	" soft herb ",
	" soft roll ",
	" soft tofu ",
	" sou vides ",
	" soup mixs ",
	" sour mixs ",
	" sourdough ",
	" soy cream ",
	" soy flour ",
	" soy milks ",
	" soy sauce ",
	" soya milk ",
	" spaetzles ",
	" spaghetti ",
	" spare rib ",
	" spareribs ",
	" spearmint ",
	" speculoos ",
	" spice mix ",
	" spice rub ",
	" spirulina ",
	" split pea ",
	" squid ink ",
	" squirrels ",
	" srirachas ",
	" star anise ",
	" star fruit ",
	" stew beef ",
	" stew meat ",
	" strawberry ",
	" stuffings ",
	" sturgeons ",
	" sub rolls ",
	" succotash ",
	" sunchokes ",
	" sunflower ",
	" sweetcorn ",
	" swordfish ",
	" taleggios ",
	" tamarillo ",
	" tamarillos ",
	" tamarinds ",
	" tangerine ",
	" tangerines ",
	" tap water ",
	" tapenades ",
	" tarragons ",
	" tater tot ",
	" tentacles ",
	" teriyakis ",
	" thai basil ",
	" thai crab ",
	" tilapiums ",
	" tiny peas ",
	" togarashi ",
	" tomatillo ",
	" tomatoess ",
	" tonka bean ",
	" topinambur ",
	" tortillas ",
	" trail mix ",
	" trout roe ",
	" tuna fish ",
	" turbinado ",
	" turmerics ",
	" tzatzikis ",
	" ugli fruit ",
	" umeboshis ",
	" urad dals ",
	" veal bone ",
	" vegemites ",
	" vegenaise ",
	" vegetable ",
	" verjuices ",
	" vermouths ",
	" vidaliums ",
	" vin santo ",
	" watercress ",
	" watermelon ",
	" wheat bun ",
	" white rum ",
	" whitefish ",
	" wild boar ",
	" wild duck ",
	" wild rice ",
	" wood chip ",
	" zinfandel ",
	" zucchinis ",
	" absinthe ",
	" acaciums ",
	" advocaat ",
	" aged rum ",
	" ahi tuna ",
	" alcohols ",
	" alfalfas ",
	" allspice ",
	" allspices ",
	" almondss ",
	" amaranth ",
	" amaranths ",
	" amaretti ",
	" amaretto ",
	" anchovys ",
	" angelicas ",
	" aniseeds ",
	" anisette ",
	" annattos ",
	" apricots ",
	" aquavits ",
	" armagnac ",
	" artichoke ",
	" arugulas ",
	" asparagus ",
	" assembly ",
	" aubergine ",
	" avocados ",
	" baby pea ",
	" baguette ",
	" baharats ",
	" balsamic ",
	" bananass ",
	" barbecue ",
	" barberry ",
	" bay leaf ",
	" bay leafs ",
	" bechamel ",
	" beef fat ",
	" beef rib ",
	" beetroot ",
	" berberes ",
	" bilberrys ",
	" biscotti ",
	" biscuits ",
	" bisquick ",
	" blueberry ",
	" bluefish ",
	" bok choy ",
	" bok choys ",
	" bolognas ",
	" bouillon ",
	" bourbons ",
	" boursins ",
	" branzino ",
	" bresaola ",
	" brioches ",
	" briskets ",
	" brittles ",
	" broccoli ",
	" broccolis ",
	" bucatini ",
	" buffalos ",
	" bulgogis ",
	" burdocks ",
	" burratum ",
	" burritos ",
	" cabbages ",
	" cachacas ",
	" cake mix ",
	" calabrese ",
	" calamari ",
	" callaloo ",
	" calvados ",
	" camparis ",
	" capicola ",
	" capsicum ",
	" caramels ",
	" caraways ",
	" cardamom ",
	" cardamoms ",
	" cardoons ",
	" carnitas ",
	" carrotss ",
	" cashewss ",
	" cassavas ",
	" cassiums ",
	" catfishs ",
	" cayennes ",
	" celeriac ",
	" celeriacs ",
	" challahs ",
	" chambord ",
	" chamomile ",
	" char siu ",
	" chayotes ",
	" cheddars ",
	" cherimoya ",
	" cherries ",
	" chervils ",
	" chestnut ",
	" chiantis ",
	" chickens ",
	" chickpea ",
	" chickpeas ",
	" chicorys ",
	" chipotle ",
	" chorizos ",
	" chowders ",
	" choy sum ",
	" chutneys ",
	" cilantro ",
	" cilantros ",
	" cinnamon ",
	" cinnamons ",
	" cobblers ",
	" cocktail ",
	" coconuts ",
	" codfishs ",
	" coleslaw ",
	" collards ",
	" compotes ",
	" cordials ",
	" corn cob ",
	" corn oil ",
	" cornmeal ",
	" courgette ",
	" couscous ",
	" crabmeat ",
	" crackers ",
	" craisins ",
	" cranberry ",
	" crawfish ",
	" criminis ",
	" crostini ",
	" croutons ",
	" cucumber ",
	" cucumbers ",
	" cucuzzas ",
	" culantro ",
	" culantros ",
	" cupcakes ",
	" cura aos ",
	" curacaos ",
	" currants ",
	" curtidos ",
	" custards ",
	" deli ham ",
	" delicatas ",
	" demerara ",
	" desserts ",
	" dill seed ",
	" ditalini ",
	" dogfishs ",
	" drambuie ",
	" dressing ",
	" dubonnet ",
	" duck egg ",
	" duck fat ",
	" duckling ",
	" dumpling ",
	" ear corn ",
	" edamames ",
	" egg roll ",
	" egg wash ",
	" egg yolk ",
	" eggplant ",
	" eggplants ",
	" emmental ",
	" empanada ",
	" epazotes ",
	" escarole ",
	" espresso ",
	" essences ",
	" falafels ",
	" falernum ",
	" farfalle ",
	" fatbacks ",
	" fenugreek ",
	" fig jams ",
	" fish oil ",
	" flaxseed ",
	" flounder ",
	" foie gra ",
	" fondants ",
	" fontinas ",
	" freekehs ",
	" fregolas ",
	" fritters ",
	" frosting ",
	" fructose ",
	" furikake ",
	" fusillis ",
	" gai lans ",
	" galangas ",
	" galangal ",
	" ganaches ",
	" garnishs ",
	" gelatins ",
	" gelatine ",
	" gemellis ",
	" genevers ",
	" gianduja ",
	" ginsengs ",
	" glucoses ",
	" gnocchis ",
	" gold rum ",
	" granolas ",
	" groupers ",
	" gruyeres ",
	" guajillo ",
	" guar gum ",
	" guinness ",
	" habanero ",
	" habaneros ",
	" haddocks ",
	" halibuts ",
	" halloumi ",
	" ham bone ",
	" ham hock ",
	" harissas ",
	" hazelnut ",
	" hibiscus ",
	" honeydew ",
	" honeydews ",
	" icebergs ",
	" irm tofu ",
	" jackfruit ",
	" jaggerys ",
	" jalapeno ",
	" jalapenos ",
	" jasmines ",
	" ketchups ",
	" key lime ",
	" kielbasa ",
	" kinh gioi ",
	" kohlrabi ",
	" kohlrabis ",
	" kool aid ",
	" korarimas ",
	" kumquats ",
	" lamb rib ",
	" lasagnas ",
	" lasagnes ",
	" lavender ",
	" lavenders ",
	" lecithin ",
	" lemon ze ",
	" lemonade ",
	" lettuces ",
	" lima bean ",
	" limeades ",
	" linguine ",
	" linguini ",
	" liqueurs ",
	" liquorice ",
	" lobsters ",
	" lollipop ",
	" macaroni ",
	" macaroon ",
	" mackerel ",
	" madeiras ",
	" mahimahi ",
	" manchego ",
	" mandarin ",
	" mandarine ",
	" mangetout ",
	" marinade ",
	" marjoram ",
	" marjorams ",
	" marsalas ",
	" marzipan ",
	" meat cut ",
	" meatball ",
	" meatloaf ",
	" meringue ",
	" mescluns ",
	" mirepoix ",
	" mirliton ",
	" mochikos ",
	" molasses ",
	" monkfish ",
	" mulberry ",
	" mulberrys ",
	" mung bean ",
	" mung dal ",
	" mushroom ",
	" mushrooms ",
	" mustards ",
	" nam plas ",
	" navy bean ",
	" nectarine ",
	" ni oises ",
	" nicoises ",
	" nopalito ",
	" nuoc nam ",
	" nut milk ",
	" nutellas ",
	" oat bran ",
	" oatmeals ",
	" octopuss ",
	" okra pod ",
	" old bays ",
	" orangess ",
	" oreganos ",
	" ostrichs ",
	" oz bacon ",
	" oz nduja ",
	" paccheri ",
	" pak choi ",
	" palm oil ",
	" pancakes ",
	" paprikas ",
	" parmesan ",
	" parsleys ",
	" parsnips ",
	" passatum ",
	" pastinas ",
	" pastrami ",
	" patty pan ",
	" pea pods ",
	" peachess ",
	" peanutss ",
	" pecorino ",
	" persimmon ",
	" pheasant ",
	" physaliss ",
	" pierogis ",
	" pig tail ",
	" pignolis ",
	" piment n ",
	" pimentos ",
	" pimenton ",
	" pimiento ",
	" pine nut ",
	" pineapple ",
	" pinenuts ",
	" plantain ",
	" plum jam ",
	" poblanos ",
	" polentum ",
	" pollocks ",
	" pompanos ",
	" popcorns ",
	" popsicle ",
	" porcinis ",
	" pork fat ",
	" pork rib ",
	" potatoes ",
	" poultrys ",
	" pralines ",
	" preserve ",
	" pretzels ",
	" prosecco ",
	" puddings ",
	" pumpkins ",
	" purslane ",
	" raclette ",
	" radicchio ",
	" raisinss ",
	" rambutans ",
	" raspberry ",
	" raviolis ",
	" recaitos ",
	" red wine ",
	" redfishs ",
	" rhubarbs ",
	" rice mix ",
	" riesling ",
	" rigatoni ",
	" risottos ",
	" ro wines ",
	" romaines ",
	" rosebuds ",
	" rosemary ",
	" rosemarys ",
	" rouilles ",
	" rutabaga ",
	" rutabagas ",
	" saffrons ",
	" salt cod ",
	" salt plu ",
	" saltines ",
	" sambucas ",
	" sandwich ",
	" sangrium ",
	" sardines ",
	" sassafra ",
	" sassafras ",
	" sausages ",
	" sauterne ",
	" scallion ",
	" scallions ",
	" scallops ",
	" schmaltz ",
	" schnapps ",
	" seafoods ",
	" seaweeds ",
	" sel gris ",
	" seltzers ",
	" semolina ",
	" serranos ",
	" shallots ",
	" shallots ",
	" shaoxing ",
	" shiitake ",
	" sirloins ",
	" smoothie ",
	" snap pea ",
	" snap peas ",
	" snappers ",
	" snow pea ",
	" sofritos ",
	" sorghums ",
	" sou vide ",
	" souffles ",
	" soup mix ",
	" sour mix ",
	" soy beans ",
	" soy milk ",
	" soybeans ",
	" soymilks ",
	" spaetzle ",
	" spanishs ",
	" sparerib ",
	" speculoo ",
	" spinachs ",
	" splendas ",
	" squirrel ",
	" sriracha ",
	" steviums ",
	" stiltons ",
	" stuffing ",
	" sturgeon ",
	" sub roll ",
	" sucanats ",
	" sultanas ",
	" sunchoke ",
	" tabascos ",
	" taleggio ",
	" tamarillo ",
	" tamarind ",
	" tamarinds ",
	" tangerine ",
	" tapenade ",
	" tapiocas ",
	" tarragon ",
	" tarragons ",
	" tentacle ",
	" tequilas ",
	" teriyaki ",
	" tilapium ",
	" tiny pea ",
	" tomatoes ",
	" tortilla ",
	" tostadas ",
	" treacles ",
	" trevisos ",
	" tri tips ",
	" trotters ",
	" truffles ",
	" truviums ",
	" turkishs ",
	" turmeric ",
	" turmerics ",
	" tzatziki ",
	" umeboshi ",
	" urad dal ",
	" vanillas ",
	" vegemite ",
	" venisons ",
	" verjuice ",
	" vermouth ",
	" vidalium ",
	" vinegars ",
	" walleyes ",
	" walnutss ",
	" whiskeys ",
	" whitings ",
	" woodruffs ",
	" yoghurts ",
	" z vodkas ",
	" za atars ",
	" zucchini ",
	" zucchinis ",
	" acacium ",
	" add ins ",
	" ajwains ",
	" alcohol ",
	" alfalfa ",
	" allspice ",
	" almonds ",
	" almonds ",
	" amaranth ",
	" amchoors ",
	" amchurs ",
	" anchovy ",
	" angelica ",
	" aniseed ",
	" annatto ",
	" aperols ",
	" appless ",
	" apricot ",
	" apricots ",
	" aquavit ",
	" arugula ",
	" arugulas ",
	" asiagos ",
	" avocado ",
	" avocados ",
	" baharat ",
	" bananas ",
	" bananas ",
	" barleys ",
	" barolos ",
	" bay leaf ",
	" berbere ",
	" bilberry ",
	" biscuit ",
	" bisques ",
	" bok choy ",
	" bologna ",
	" bonitos ",
	" bourbon ",
	" boursin ",
	" brandys ",
	" brioche ",
	" brisket ",
	" brittle ",
	" broccoli ",
	" brownys ",
	" buffalo ",
	" bulgogi ",
	" bulgurs ",
	" burdock ",
	" burgers ",
	" burrito ",
	" butters ",
	" cabbage ",
	" cabbages ",
	" cachaca ",
	" calvado ",
	" campari ",
	" camphors ",
	" canolas ",
	" cantals ",
	" caperss ",
	" caramel ",
	" caraway ",
	" caraways ",
	" cardamom ",
	" cardoon ",
	" carnita ",
	" carrots ",
	" carrots ",
	" cashews ",
	" cashews ",
	" cassava ",
	" cassium ",
	" catfish ",
	" catsups ",
	" caviars ",
	" cayenne ",
	" celeriac ",
	" celerys ",
	" cereals ",
	" ch vres ",
	" challah ",
	" charolis ",
	" chayote ",
	" chayotes ",
	" cheddar ",
	" cheeses ",
	" cherrys ",
	" chervil ",
	" chervils ",
	" chianti ",
	" chicken ",
	" chickpea ",
	" chicory ",
	" chivess ",
	" chorizo ",
	" chowder ",
	" chutney ",
	" cilantro ",
	" cinnamon ",
	" citrons ",
	" cobbler ",
	" cockles ",
	" coconut ",
	" coconuts ",
	" codfish ",
	" coffees ",
	" cognacs ",
	" collard ",
	" comices ",
	" compote ",
	" confits ",
	" cookies ",
	" cordial ",
	" cotijas ",
	" couscou ",
	" cracker ",
	" craisin ",
	" creamys ",
	" crimini ",
	" criscos ",
	" crouton ",
	" cucumber ",
	" cucuzza ",
	" culantro ",
	" cupcake ",
	" cura ao ",
	" curacao ",
	" currant ",
	" currants ",
	" curtido ",
	" custard ",
	" daikons ",
	" delicata ",
	" dessert ",
	" dogfish ",
	" dragees ",
	" dukkahs ",
	" edamame ",
	" eggnogs ",
	" eggplant ",
	" endifes ",
	" endives ",
	" epazote ",
	" epazotes ",
	" essence ",
	" fajitas ",
	" falafel ",
	" farinas ",
	" fatback ",
	" fennels ",
	" fig jam ",
	" fondant ",
	" fondues ",
	" fontina ",
	" freekeh ",
	" fregola ",
	" frenchs ",
	" frescas ",
	" frisees ",
	" fritter ",
	" fusilli ",
	" gai lan ",
	" galanga ",
	" ganache ",
	" garlics ",
	" garnish ",
	" gelatin ",
	" gelatos ",
	" gemelli ",
	" genever ",
	" gingers ",
	" ginseng ",
	" glucose ",
	" glutens ",
	" gnocchi ",
	" granola ",
	" grappas ",
	" gratins ",
	" greases ",
	" grouper ",
	" gruyere ",
	" habanero ",
	" haddock ",
	" halibut ",
	" halvahs ",
	" harissa ",
	" hibiscu ",
	" hijikis ",
	" hoisins ",
	" hominys ",
	" honeydew ",
	" iceberg ",
	" injeras ",
	" jaggery ",
	" jalapeno ",
	" jasmine ",
	" jicamas ",
	" kahluas ",
	" kaisers ",
	" kernels ",
	" ketchup ",
	" kimchis ",
	" kirschs ",
	" kohlrabi ",
	" korarima ",
	" kumaras ",
	" kumquat ",
	" kumquats ",
	" labnehs ",
	" lardons ",
	" lasagna ",
	" lasagne ",
	" laurels ",
	" lavashs ",
	" lavender ",
	" lemonss ",
	" lentils ",
	" lentilss ",
	" lettuce ",
	" lettuces ",
	" lillets ",
	" limeade ",
	" linguis ",
	" liqueur ",
	" liquors ",
	" lobster ",
	" loquats ",
	" lovages ",
	" lychees ",
	" madeira ",
	" maldons ",
	" marjoram ",
	" marsala ",
	" masalas ",
	" mastics ",
	" matchas ",
	" matzohs ",
	" mesclun ",
	" mezcals ",
	" midoris ",
	" millets ",
	" mitsubas ",
	" mizunas ",
	" mochiko ",
	" mousses ",
	" mueslis ",
	" muffins ",
	" mugworts ",
	" mulberry ",
	" mushroom ",
	" mussels ",
	" mustard ",
	" muttons ",
	" myrtles ",
	" nam pla ",
	" nectars ",
	" ni oise ",
	" nicoise ",
	" nigellas ",
	" njangsas ",
	" noodles ",
	" nopales ",
	" nougats ",
	" noyauxs ",
	" nutella ",
	" nutmegs ",
	" oatmeal ",
	" octopus ",
	" old bay ",
	" olivess ",
	" onionss ",
	" oranges ",
	" oranges ",
	" orchids ",
	" oregano ",
	" oreganos ",
	" orgeats ",
	" ostrich ",
	" oxtails ",
	" oysters ",
	" pancake ",
	" paneers ",
	" panelas ",
	" papayas ",
	" paprika ",
	" paprikas ",
	" parsley ",
	" parsleys ",
	" parsnip ",
	" parsnips ",
	" pastina ",
	" pastrys ",
	" pawpaws ",
	" pea pod ",
	" peaches ",
	" peanuts ",
	" peanuts ",
	" pecanss ",
	" pectins ",
	" pepitas ",
	" peppers ",
	" perillas ",
	" pernods ",
	" phyllos ",
	" physalis ",
	" pickles ",
	" pierogi ",
	" pignoli ",
	" pimento ",
	" pinenut ",
	" pistous ",
	" poblano ",
	" pollock ",
	" pomelos ",
	" pompano ",
	" popcorn ",
	" porcini ",
	" potatos ",
	" poultry ",
	" praline ",
	" pretzel ",
	" pudding ",
	" pumpkin ",
	" pumpkins ",
	" quahogs ",
	" quiches ",
	" quinces ",
	" quinoas ",
	" rabbits ",
	" radhunis ",
	" radishs ",
	" raisins ",
	" raisins ",
	" rambutan ",
	" rapinis ",
	" ravioli ",
	" recaito ",
	" redfish ",
	" relishs ",
	" rennets ",
	" rhubarb ",
	" rhubarbs ",
	" ribeyes ",
	" risotto ",
	" ro wine ",
	" rockets ",
	" romaine ",
	" rooibos ",
	" rosebud ",
	" rosemary ",
	" rotinis ",
	" rouille ",
	" russets ",
	" rutabaga ",
	" saffron ",
	" saffrons ",
	" salamis ",
	" salmons ",
	" saltine ",
	" sambals ",
	" sambuca ",
	" sardine ",
	" satsumas ",
	" sausage ",
	" savorys ",
	" scallion ",
	" scallop ",
	" scampis ",
	" schnapp ",
	" scotchs ",
	" seafood ",
	" seaweed ",
	" seitans ",
	" sel gri ",
	" seltzer ",
	" serrano ",
	" sesames ",
	" shallot ",
	" shallots ",
	" sherrys ",
	" shrimps ",
	" sirloin ",
	" skewers ",
	" skirrets ",
	" snap pea ",
	" snapper ",
	" sofrito ",
	" sorbets ",
	" sorghum ",
	" sorrels ",
	" souffle ",
	" soy bean ",
	" soybean ",
	" soymilk ",
	" spanish ",
	" spinach ",
	" spinachs ",
	" spirits ",
	" splenda ",
	" spreads ",
	" sprites ",
	" sprouts ",
	" squashs ",
	" squash s ",
	" starchs ",
	" stevium ",
	" stilton ",
	" sucanat ",
	" sultana ",
	" tabasco ",
	" tahinis ",
	" tamales ",
	" tamaris ",
	" tamarind ",
	" tapioca ",
	" taramas ",
	" tarragon ",
	" tat sois ",
	" tempehs ",
	" tequila ",
	" tobikos ",
	" toffees ",
	" tomatos ",
	" tostada ",
	" treacle ",
	" treviso ",
	" tri tip ",
	" trotter ",
	" truffle ",
	" truvium ",
	" turbots ",
	" turkeys ",
	" turkish ",
	" turmeric ",
	" turnips ",
	" vanilla ",
	" vanillas ",
	" venison ",
	" vinegar ",
	" waffles ",
	" wakames ",
	" walleye ",
	" walnuts ",
	" walnuts ",
	" wasabis ",
	" whiskey ",
	" whiskys ",
	" whiting ",
	" wontons ",
	" woodruff ",
	" yoghurt ",
	" yogurts ",
	" z vodka ",
	" za atar ",
	" zedoarys ",
	" zereshks ",
	" zucchini ",
	" add in ",
	" adobos ",
	" agaves ",
	" aiolis ",
	" ajwain ",
	" ajwains ",
	" almond ",
	" amaros ",
	" ambers ",
	" amchoor ",
	" amchur ",
	" anchos ",
	" anises ",
	" aonoris ",
	" aperol ",
	" apples ",
	" apples ",
	" apricot ",
	" arames ",
	" arugula ",
	" asiago ",
	" avocado ",
	" bacons ",
	" bagels ",
	" banana ",
	" bananas ",
	" barley ",
	" barolo ",
	" basils ",
	" beanss ",
	" beetss ",
	" berrys ",
	" bisons ",
	" bisque ",
	" bonito ",
	" borages ",
	" brandy ",
	" breads ",
	" breams ",
	" brines ",
	" broths ",
	" browny ",
	" bugles ",
	" bulgur ",
	" burger ",
	" butter ",
	" cabbage ",
	" cacaos ",
	" cachas ",
	" cactus ",
	" cajuns ",
	" camphor ",
	" candys ",
	" canola ",
	" cantal ",
	" capers ",
	" capers ",
	" caraway ",
	" carobs ",
	" carrot ",
	" carrots ",
	" cashew ",
	" cassis ",
	" cassias ",
	" catsup ",
	" caviar ",
	" cedars ",
	" celery ",
	" celerys ",
	" cereal ",
	" ch vre ",
	" chards ",
	" charoli ",
	" chayote ",
	" cheese ",
	" chenpis ",
	" cherry ",
	" cherrys ",
	" chervil ",
	" chiles ",
	" chilis ",
	" chives ",
	" chives ",
	" chivess ",
	" chucks ",
	" cicelys ",
	" ciders ",
	" citron ",
	" citrus ",
	" cloves ",
	" cockle ",
	" cocoas ",
	" coconut ",
	" coffee ",
	" cognac ",
	" colbys ",
	" comice ",
	" conchs ",
	" confit ",
	" cookie ",
	" cophas ",
	" coppas ",
	" cotija ",
	" creams ",
	" creamy ",
	" cremas ",
	" crepes ",
	" crisco ",
	" crisps ",
	" crusts ",
	" cumins ",
	" currant ",
	" currys ",
	" cynars ",
	" daikon ",
	" daikons ",
	" damsons ",
	" dashis ",
	" datess ",
	" dijons ",
	" donuts ",
	" doughs ",
	" dragee ",
	" drinks ",
	" dukkah ",
	" durians ",
	" eggnog ",
	" endife ",
	" endive ",
	" endives ",
	" epazote ",
	" equals ",
	" fajita ",
	" farina ",
	" farros ",
	" feijoas ",
	" fennel ",
	" fennels ",
	" fetums ",
	" fideos ",
	" filets ",
	" flours ",
	" flukes ",
	" fondue ",
	" franks ",
	" french ",
	" fresca ",
	" fri es ",
	" frisee ",
	" frisees ",
	" fritos ",
	" fruits ",
	" fudges ",
	" garlic ",
	" garlics ",
	" gelato ",
	" ginger ",
	" gingers ",
	" gluten ",
	" golpars ",
	" gooses ",
	" goudas ",
	" grapes ",
	" grappa ",
	" grasss ",
	" gratin ",
	" gravys ",
	" grease ",
	" greeks ",
	" greens ",
	" guavas ",
	" gumbos ",
	" halvah ",
	" herbes ",
	" hijiki ",
	" hoisin ",
	" hominy ",
	" honeys ",
	" hummus ",
	" hyssops ",
	" icings ",
	" injera ",
	" jambuls ",
	" jellos ",
	" jellys ",
	" jicama ",
	" jicamas ",
	" juices ",
	" jujubes ",
	" kabobs ",
	" kahlua ",
	" kaiser ",
	" kamuts ",
	" kashas ",
	" kebabs ",
	" kefirs ",
	" kernel ",
	" kimchi ",
	" kirsch ",
	" kombus ",
	" konbus ",
	" kumara ",
	" kumquat ",
	" labneh ",
	" lagers ",
	" lardos ",
	" lardon ",
	" laurel ",
	" lavash ",
	" leekss ",
	" legumes ",
	" lemons ",
	" lemons ",
	" lentil ",
	" lentils ",
	" lettuce ",
	" licors ",
	" lillet ",
	" lingui ",
	" liquor ",
	" livers ",
	" loafes ",
	" loquat ",
	" loquats ",
	" lovage ",
	" lovages ",
	" lychee ",
	" lychees ",
	" m ches ",
	" mahlabs ",
	" maldon ",
	" mangos ",
	" maples ",
	" masala ",
	" mastic ",
	" mastics ",
	" matcha ",
	" matzos ",
	" matzoh ",
	" melons ",
	" mezcal ",
	" midori ",
	" millet ",
	" mirins ",
	" mitsuba ",
	" mizuna ",
	" mooses ",
	" morels ",
	" mousse ",
	" muesli ",
	" muffin ",
	" mugwort ",
	" mussel ",
	" mutton ",
	" myrtle ",
	" nachos ",
	" ndujas ",
	" nectar ",
	" nigella ",
	" njangsa ",
	" noodle ",
	" nopals ",
	" nopale ",
	" nopales ",
	" nougat ",
	" noyaux ",
	" nutmeg ",
	" nutmegs ",
	" olives ",
	" olives ",
	" onions ",
	" onions ",
	" orange ",
	" oranges ",
	" orchid ",
	" oregano ",
	" orgeat ",
	" oxtail ",
	" oyster ",
	" pamelos ",
	" paneer ",
	" panela ",
	" panirs ",
	" pankos ",
	" pansys ",
	" papaya ",
	" papayas ",
	" paprika ",
	" parsley ",
	" parsnip ",
	" pastas ",
	" pastis ",
	" pastry ",
	" pawpaw ",
	" peachs ",
	" peanut ",
	" pearss ",
	" pecans ",
	" pecans ",
	" pectin ",
	" pennes ",
	" pepita ",
	" pepper ",
	" perchs ",
	" perilla ",
	" pernod ",
	" pestos ",
	" phyllo ",
	" pickle ",
	" pilafs ",
	" piscos ",
	" pistou ",
	" pitums ",
	" pizzas ",
	" pluots ",
	" pomelo ",
	" pomelos ",
	" ponzus ",
	" potato ",
	" potatos ",
	" prawns ",
	" prunes ",
	" pumpkin ",
	" quahog ",
	" quails ",
	" quarks ",
	" quesos ",
	" quiche ",
	" quince ",
	" quinces ",
	" quinoa ",
	" rabbit ",
	" radhuni ",
	" radish ",
	" radishs ",
	" raisin ",
	" raisins ",
	" ramans ",
	" rapini ",
	" relish ",
	" rennet ",
	" rhubarb ",
	" ribeye ",
	" roasts ",
	" rocket ",
	" rooibo ",
	" rotels ",
	" rotini ",
	" russet ",
	" saffron ",
	" salads ",
	" salami ",
	" salmon ",
	" salsas ",
	" sambal ",
	" sanshos ",
	" satsuma ",
	" sauces ",
	" savory ",
	" savorys ",
	" scampi ",
	" scones ",
	" scotch ",
	" scrods ",
	" seitan ",
	" sesame ",
	" sesames ",
	" shallot ",
	" sharks ",
	" sheeps ",
	" sherry ",
	" shisos ",
	" shoyus ",
	" shrimp ",
	" skewer ",
	" skirret ",
	" smelts ",
	" snails ",
	" sorbet ",
	" sorrel ",
	" sorrels ",
	" specks ",
	" spelts ",
	" spinach ",
	" spirit ",
	" sprays ",
	" spread ",
	" sprite ",
	" sprout ",
	" squash ",
	" squash ",
	" squids ",
	" starch ",
	" steaks ",
	" stouts ",
	" sugars ",
	" sumacs ",
	" sumaqs ",
	" swedes ",
	" syrups ",
	" tahini ",
	" tamale ",
	" tamari ",
	" tarama ",
	" tassos ",
	" tat soi ",
	" tempeh ",
	" thymes ",
	" toasts ",
	" tobiko ",
	" toffee ",
	" tomato ",
	" tortes ",
	" tripes ",
	" trouts ",
	" tuberss ",
	" turbot ",
	" turkey ",
	" turnip ",
	" turnips ",
	" vanilla ",
	" verjus ",
	" vodkas ",
	" wafers ",
	" waffle ",
	" wakame ",
	" walnut ",
	" wasabi ",
	" wasabis ",
	" waters ",
	" wheats ",
	" whisky ",
	" wonton ",
	" yeasts ",
	" yogurt ",
	" yuccas ",
	" zedoary ",
	" zereshk ",
	" adobo ",
	" agars ",
	" agave ",
	" aioli ",
	" ajwain ",
	" alums ",
	" amaro ",
	" amber ",
	" ancho ",
	" anise ",
	" anises ",
	" aonori ",
	" apple ",
	" apples ",
	" arame ",
	" bacon ",
	" bagel ",
	" banana ",
	" barks ",
	" basil ",
	" basils ",
	" basss ",
	" beans ",
	" beans ",
	" beefs ",
	" beers ",
	" beets ",
	" beets ",
	" berry ",
	" bings ",
	" bison ",
	" boars ",
	" boldos ",
	" bones ",
	" borage ",
	" brans ",
	" bread ",
	" bream ",
	" bries ",
	" brine ",
	" broth ",
	" bugle ",
	" cacao ",
	" cacha ",
	" cactu ",
	" cajun ",
	" cakes ",
	" candy ",
	" caper ",
	" carob ",
	" carps ",
	" carrot ",
	" cassi ",
	" cassia ",
	" cavas ",
	" cedar ",
	" celery ",
	" chais ",
	" chard ",
	" chards ",
	" chees ",
	" chenpi ",
	" cherry ",
	" chile ",
	" chili ",
	" chips ",
	" chive ",
	" chives ",
	" chuck ",
	" cicely ",
	" cider ",
	" citru ",
	" clams ",
	" clovs ",
	" clove ",
	" cloves ",
	" cocoa ",
	" colas ",
	" colby ",
	" conch ",
	" copha ",
	" coppa ",
	" corns ",
	" cr ms ",
	" crabs ",
	" cream ",
	" crema ",
	" crepe ",
	" cresss ",
	" crisp ",
	" crust ",
	" cubebs ",
	" cumin ",
	" cumins ",
	" curds ",
	" curry ",
	" cynar ",
	" daikon ",
	" damson ",
	" dashi ",
	" dates ",
	" dates ",
	" dijon ",
	" dills ",
	" donut ",
	" dough ",
	" drink ",
	" ducks ",
	" durian ",
	" edams ",
	" eggss ",
	" endive ",
	" equal ",
	" farro ",
	" feijoa ",
	" fennel ",
	" fetum ",
	" fideo ",
	" filet ",
	" fishs ",
	" flaxs ",
	" flour ",
	" fluke ",
	" frank ",
	" fri e ",
	" frisee ",
	" frito ",
	" fruit ",
	" fudge ",
	" garlic ",
	" ghees ",
	" ginger ",
	" goats ",
	" golds ",
	" golpar ",
	" goose ",
	" gouda ",
	" grape ",
	" grapes ",
	" grass ",
	" gravy ",
	" greek ",
	" green ",
	" grits ",
	" gruys ",
	" guava ",
	" guavas ",
	" gumbo ",
	" hakes ",
	" hashs ",
	" herbs ",
	" herbe ",
	" hings ",
	" honey ",
	" hummu ",
	" hyssop ",
	" icing ",
	" jambul ",
	" jello ",
	" jelly ",
	" jerks ",
	" jicama ",
	" jimbus ",
	" juice ",
	" jujube ",
	" kabob ",
	" kahls ",
	" kales ",
	" kamut ",
	" kasha ",
	" kebab ",
	" kefir ",
	" kelps ",
	" kiwis ",
	" kokums ",
	" kombu ",
	" konbu ",
	" lager ",
	" lambs ",
	" lards ",
	" lardo ",
	" leeks ",
	" leek s ",
	" leeks ",
	" legume ",
	" lemon ",
	" lemons ",
	" licor ",
	" limes ",
	" liver ",
	" loafe ",
	" loquat ",
	" lotus ",
	" lovage ",
	" lychee ",
	" m che ",
	" maces ",
	" mahlab ",
	" malts ",
	" mameys ",
	" mango ",
	" mangos ",
	" maple ",
	" masas ",
	" mastic ",
	" mates ",
	" matzo ",
	" meats ",
	" melon ",
	" milks ",
	" mints ",
	" mirin ",
	" misos ",
	" moose ",
	" morel ",
	" naans ",
	" nacho ",
	" nduja ",
	" nopal ",
	" nopale ",
	" noris ",
	" nutmeg ",
	" nutss ",
	" oatss ",
	" okras ",
	" oleos ",
	" olive ",
	" olives ",
	" onion ",
	" onions ",
	" orange ",
	" oreos ",
	" orzos ",
	" ouzos ",
	" palms ",
	" pamelo ",
	" panir ",
	" panko ",
	" pansy ",
	" papaya ",
	" pasta ",
	" pasti ",
	" pates ",
	" peach ",
	" peachs ",
	" pears ",
	" pears ",
	" peass ",
	" pecan ",
	" penne ",
	" perch ",
	" pesto ",
	" pilaf ",
	" pimms ",
	" pines ",
	" pisco ",
	" pitas ",
	" pitum ",
	" pizza ",
	" plums ",
	" pluot ",
	" pomelo ",
	" ponzu ",
	" porks ",
	" ports ",
	" potato ",
	" prawn ",
	" prune ",
	" quail ",
	" quark ",
	" queso ",
	" quince ",
	" radish ",
	" ragus ",
	" raisin ",
	" raman ",
	" ramps ",
	" rices ",
	" roast ",
	" roses ",
	" rotel ",
	" rotis ",
	" rumps ",
	" sabas ",
	" sages ",
	" sagos ",
	" sakes ",
	" salad ",
	" salsa ",
	" salts ",
	" sansho ",
	" sauce ",
	" savory ",
	" scone ",
	" scrod ",
	" seeds ",
	" sesame ",
	" shark ",
	" sheep ",
	" shiso ",
	" shisos ",
	" shoyu ",
	" skins ",
	" smelt ",
	" snail ",
	" sobas ",
	" sodas ",
	" sojus ",
	" soles ",
	" sorrel ",
	" soups ",
	" spams ",
	" speck ",
	" spelt ",
	" spray ",
	" squid ",
	" steak ",
	" stews ",
	" stout ",
	" suets ",
	" sugar ",
	" sumac ",
	" sumacs ",
	" sumaq ",
	" swede ",
	" syrup ",
	" tacos ",
	" taros ",
	" tarts ",
	" tasso ",
	" teffs ",
	" thrus ",
	" thyme ",
	" thymes ",
	" toast ",
	" tofus ",
	" torte ",
	" tripe ",
	" trout ",
	" tubers ",
	" tunas ",
	" turnip ",
	" udons ",
	" urfas ",
	" uzazis ",
	" veals ",
	" verju ",
	" vodka ",
	" wafer ",
	" wasabi ",
	" water ",
	" wheat ",
	" wheys ",
	" wines ",
	" wraps ",
	" yeast ",
	" yolks ",
	" yucas ",
	" yucca ",
	" zitis ",
	" agar ",
	" ahis ",
	" ales ",
	" alum ",
	" ames ",
	" anise ",
	" apple ",
	" bark ",
	" basil ",
	" bass ",
	" bays ",
	" bean ",
	" beef ",
	" beer ",
	" beet ",
	" beets ",
	" bing ",
	" boar ",
	" boldo ",
	" bone ",
	" bran ",
	" brie ",
	" buns ",
	" cake ",
	" cals ",
	" carp ",
	" cava ",
	" chai ",
	" chard ",
	" chee ",
	" chip ",
	" clam ",
	" clov ",
	" clove ",
	" cods ",
	" cola ",
	" corn ",
	" corns ",
	" cr m ",
	" crab ",
	" cress ",
	" cubeb ",
	" cumin ",
	" curd ",
	" date ",
	" dates ",
	" dill ",
	" dills ",
	" duck ",
	" edam ",
	" eels ",
	" eggs ",
	" eggs ",
	" fats ",
	" figs ",
	" fish ",
	" flax ",
	" ghee ",
	" gins ",
	" goat ",
	" gold ",
	" grape ",
	" grit ",
	" gruy ",
	" guava ",
	" hake ",
	" hams ",
	" hash ",
	" hemps ",
	" herb ",
	" hing ",
	" hops ",
	" ices ",
	" jams ",
	" jerk ",
	" jimbu ",
	" kahl ",
	" kale ",
	" kales ",
	" kelp ",
	" kiwi ",
	" kokum ",
	" lamb ",
	" lard ",
	" leek ",
	" leek ",
	" lemon ",
	" lime ",
	" limes ",
	" lotu ",
	" loxs ",
	" m ms ",
	" mace ",
	" maces ",
	" malt ",
	" mamey ",
	" mango ",
	" masa ",
	" mate ",
	" meat ",
	" milk ",
	" mint ",
	" mints ",
	" miso ",
	" msgs ",
	" naan ",
	" nori ",
	" nuts ",
	" nuts ",
	" oats ",
	" oats ",
	" oils ",
	" okra ",
	" okras ",
	" oleo ",
	" olive ",
	" onion ",
	" oreo ",
	" orzo ",
	" ouzo ",
	" palm ",
	" pate ",
	" peas ",
	" peach ",
	" pear ",
	" pears ",
	" peas ",
	" pies ",
	" pimm ",
	" pine ",
	" pita ",
	" plus ",
	" plum ",
	" plums ",
	" pork ",
	" port ",
	" ragu ",
	" ramp ",
	" ribs ",
	" rice ",
	" roes ",
	" rose ",
	" roses ",
	" roti ",
	" rubs ",
	" rums ",
	" rump ",
	" ryes ",
	" saba ",
	" sage ",
	" sages ",
	" sago ",
	" sake ",
	" salt ",
	" salts ",
	" seed ",
	" shiso ",
	" skin ",
	" soba ",
	" soda ",
	" soju ",
	" sole ",
	" soup ",
	" spam ",
	" stew ",
	" suet ",
	" sumac ",
	" taco ",
	" taro ",
	" taros ",
	" tart ",
	" teff ",
	" thru ",
	" thyme ",
	" tips ",
	" tofu ",
	" tuna ",
	" udon ",
	" urfa ",
	" uzazi ",
	" veal ",
	" whey ",
	" wine ",
	" wrap ",
	" yams ",
	" yolk ",
	" yuca ",
	" yuzus ",
	" zests ",
	" ziti ",
	" ahi ",
	" ale ",
	" ame ",
	" bay ",
	" beet ",
	" bun ",
	" cal ",
	" cod ",
	" corn ",
	" date ",
	" dill ",
	" eel ",
	" egg ",
	" fat ",
	" fig ",
	" figs ",
	" gin ",
	" ham ",
	" hemp ",
	" hop ",
	" ice ",
	" jam ",
	" kale ",
	" lime ",
	" lox ",
	" m m ",
	" mace ",
	" mint ",
	" mls ",
	" msg ",
	" nut ",
	" nuts ",
	" oat ",
	" oil ",
	" okra ",
	" pea ",
	" pear ",
	" peas ",
	" pie ",
	" plu ",
	" plum ",
	" rib ",
	" ros ",
	" roe ",
	" rose ",
	" rub ",
	" rues ",
	" rum ",
	" rye ",
	" sage ",
	" salt ",
	" taro ",
	" tip ",
	" yam ",
	" yams ",
	" yuzu ",
	" zest ",
	" fig ",
	" ml ",
	" nut ",
	" pea ",
	" ro ",
	" rue ",
	" yam "}
var corpusMeasures = []string{" tablespoons ",
	" milliliter ",
	" tablespoon ",
	" teaspoons ",
	" teaspoon ",
	" canned ",
	" ounces ",
	" pounds ",
	" quarts ",
	" grams ",
	" ounce ",
	" pints ",
	" pound ",
	" quart ",
	" tbsps ",
	" cans ",
	" cups ",
	" gram ",
	" pint ",
	" tbsp ",
	" tsps ",
	" can ",
	" cup ",
	" tbl ",
	" tsp ",
	" ml ",
	" oz ",
	" c ",
	" g "}
var corpusNumbers = []string{" 1/2 ",
	" 1/3 ",
	" 1/4 ",
	" 1/5 ",
	" 1/6 ",
	" 1/7 ",
	" 1/8 ",
	" 2/3 ",
	" 2/5 ",
	" 2/7 ",
	" 3/4 ",
	" 3/8 ",
	" 4/5 ",
	" 5/8 ",
	" 7/8 ",
	" 1 ",
	" 2 ",
	" 3 ",
	" 4 ",
	" 5 ",
	" 6 ",
	" 7 ",
	" 8 ",
	" 9 ",
	" 10 ",
	" 11 ",
	" 12 ",
	" 13 ",
	" 14 ",
	" 15 ",
	" 16 ",
	" 17 ",
	" 18 ",
	" 19 ",
	" 20 ",
	" ⅜ ",
	" ⅞ ",
	" ⅔ ",
	" ⅝ ",
	" ⅓ ",
	" ½ ",
	" ¼ ",
	" ¾ ",
	" ⅛ "}
var corpusDirections = []string{" 1 ",
	" 10 ",
	" 15 ",
	" 2 ",
	" 25 ",
	" 375 ",
	" 6 ",
	" 8 ",
	" 9x13 ",
	" a ",
	" about ",
	" achieve ",
	" add ",
	" additional ",
	" an ",
	" and ",
	" arrange ",
	" assemble ",
	" bake ",
	" baking ",
	" basil ",
	" beef ",
	" before ",
	" boil ",
	" boiling ",
	" bottom ",
	" bowl ",
	" bring ",
	" brown ",
	" browned ",
	" butter ",
	" cheese ",
	" chopped ",
	" coconut ",
	" cold ",
	" color ",
	" combine ",
	" completely ",
	" cook ",
	" cookie ",
	" cooking ",
	" cool ",
	" corn ",
	" cover ",
	" covered ",
	" crushed ",
	" cup ",
	" cups ",
	" degrees ",
	" dish ",
	" distributed ",
	" does ",
	" drain ",
	" dutch ",
	" eggs ",
	" either ",
	" even ",
	" evenly ",
	" everything ",
	" f ",
	" fennel ",
	" foil ",
	" for ",
	" from ",
	" garlic ",
	" grab ",
	" ground ",
	" half ",
	" heat ",
	" hours ",
	" in ",
	" inch ",
	" into ",
	" is ",
	" italian ",
	" large ",
	" lasagna ",
	" layers ",
	" lengthwise ",
	" let ",
	" lightly ",
	" make ",
	" meat ",
	" medium ",
	" melted ",
	" minutes ",
	" mix ",
	" mixing ",
	" mixture ",
	" mozzarella ",
	" next ",
	" noodles ",
	" not ",
	" nutmeg ",
	" nuts ",
	" oats ",
	" occasionally ",
	" of ",
	" on ",
	" one ",
	" onion ",
	" or ",
	" oven ",
	" over ",
	" pans ",
	" parmesan ",
	" parsley ",
	" paste ",
	" pecans ",
	" pepper ",
	" pie ",
	" place ",
	" pot ",
	" pour ",
	" preheat ",
	" preheated ",
	" prevent ",
	" put ",
	" raisins ",
	" remaining ",
	" remove ",
	" repeat ",
	" ricotta ",
	" rinse ",
	" salt ",
	" salted ",
	" sauce ",
	" sausage ",
	" season ",
	" seasoning ",
	" seeds ",
	" serving ",
	" sheet ",
	" shell ",
	" simmer ",
	" slices ",
	" spoon ",
	" spray ",
	" spread ",
	" sprinkle ",
	" sticking ",
	" stir ",
	" stirring ",
	" sugar ",
	" sure ",
	" syrup ",
	" tablespoon ",
	" tablespoons ",
	" teaspoon ",
	" the ",
	" then ",
	" third ",
	" to ",
	" tomato ",
	" tomatoes ",
	" top ",
	" touch ",
	" transfer ",
	" until ",
	" water ",
	" well ",
	" with "}
var corpusDirectionsNeg = []string{" a ",
	" all ",
	" also ",
	" amount ",
	" and ",
	" be ",
	" best ",
	" brand ",
	" butter ",
	" c ",
	" calcium ",
	" calories ",
	" carbohydrates ",
	" chip ",
	" chocolate ",
	" cholesterol ",
	" cookie ",
	" cookies. ",
	" costco ",
	" dessert ",
	" dinner ",
	" dough ",
	" ensure ",
	" equally ",
	" ever ",
	" excellent ",
	" fat ",
	" fiber ",
	" for ",
	" from ",
	" great! ",
	" have ",
	" i ",
	" images ",
	" iron ",
	" it's ",
	" just ",
	" kirkland ",
	" liking! ",
	" make ",
	" per ",
	" potassium ",
	" protein ",
	" recipe ",
	" recommend ",
	" results. ",
	" salted ",
	" saturated ",
	" serving ",
	" servings ",
	" sodium ",
	" tasting ",
	" text ",
	" that ",
	" the ",
	" then ",
	" these ",
	" tillamook ",
	" to ",
	" unsalted ",
	" use ",
	" used ",
	" vitamin ",
	" with ",
	" would ",
	" yield ",
	" your "}

type fractionNumber struct {
	fractionString string
	value          float64
}

var corpusFractionNumberMap = map[string]fractionNumber{
	"¼": {"1/4", 0.2500000000},
	"½": {"1/2", 0.5000000000},
	"¾": {"3/4", 0.7500000000},
	"⅓": {"1/3", 0.3333333333},
	"⅔": {"2/3", 0.6666666667},
	"⅛": {"1/8", 0.1250000000},
	"⅜": {"3/8", 0.3750000000},
	"⅝": {"5/8", 0.6250000000},
	"⅞": {"7/8", 0.8750000000},
}

var corpusMeasuresMap = map[string]string{
	"c":           "cup",
	"can":         "can",
	"canned":      "can",
	"cans":        "can",
	"cup":         "cup",
	"cups":        "cup",
	"g":           "gram",
	"gram":        "gram",
	"grams":       "gram",
	"milliliter":  "milliliter",
	"ml":          "milliliter",
	"ounce":       "ounce",
	"ounces":      "ounce",
	"oz":          "ounce",
	"pint":        "pint",
	"pints":       "pint",
	"pound":       "pound",
	"pounds":      "pound",
	"quart":       "quart",
	"quarts":      "quart",
	"tablespoon":  "tbl",
	"tablespoons": "tbl",
	"tbl":         "tbl",
	"tbsp":        "tbl",
	"tbsps":       "tbl",
	"teaspoon":    "tsp",
	"teaspoons":   "tsp",
	"tsp":         "tsp",
	"tsps":        "tsp",
}

var densities = map[string]float64{
	"almond milk":            245.5000000000,
	"almonds":                115.1000000000,
	"apple":                  144.6000000000,
	"apple juice":            243.5000000000,
	"apples":                 152.2000000000,
	"applesauce":             247.2000000000,
	"arugula":                20.0000000000,
	"asparagus":              204.3000000000,
	"avocado":                218.0000000000,
	"bacon":                  156.8000000000,
	"balsamic vinegar":       255.0000000000,
	"banana":                 159.1000000000,
	"bananas":                158.3000000000,
	"barbecue sauce":         279.2000000000,
	"basil":                  191.0000000000,
	"beans":                  189.5000000000,
	"beef":                   212.8000000000,
	"beef broth":             225.2000000000,
	"beef stock":             245.3000000000,
	"beer":                   237.0000000000,
	"beets":                  195.8000000000,
	"black beans":            194.5000000000,
	"blue cheese":            135.0000000000,
	"blueberries":            206.5000000000,
	"bread":                  145.8000000000,
	"bread crumbs":           114.0000000000,
	"broccoli":               154.6000000000,
	"brown rice":             174.2000000000,
	"brown sugar":            160.9000000000,
	"butter":                 227.0000000000,
	"buttermilk":             245.0000000000,
	"cabbage":                112.1000000000,
	"canola oil":             180.5000000000,
	"carrot":                 155.0000000000,
	"carrots":                178.6000000000,
	"catsup":                 240.0000000000,
	"cauliflower":            142.2000000000,
	"celery":                 182.8000000000,
	"cheddar cheese":         179.0000000000,
	"cheese":                 144.7000000000,
	"cherries":               208.2000000000,
	"chicken":                194.0000000000,
	"chicken broth":          203.2000000000,
	"chicken soup":           248.3000000000,
	"chicken stock":          240.0000000000,
	"chickpeas":              193.3000000000,
	"chile":                  37.0000000000,
	"chili sauce":            259.0000000000,
	"chives":                 3.2000000000,
	"chocolate":              171.2000000000,
	"cider vinegar":          239.0000000000,
	"cinnamon":               53.5000000000,
	"cocoa":                  51.4000000000,
	"coconut":                261.9000000000,
	"coconut milk":           236.5000000000,
	"coconut oil":            218.0000000000,
	"coffee":                 248.7000000000,
	"coriander":              16.0000000000,
	"corn":                   164.1000000000,
	"corn flour":             123.4000000000,
	"cornmeal":               144.7000000000,
	"cornstarch":             128.0000000000,
	"cottage cheese":         216.3000000000,
	"cranberries":            123.3000000000,
	"cream":                  187.8000000000,
	"cream cheese":           237.3000000000,
	"cream of mushroom soup": 249.9000000000,
	"cucumber":               146.3000000000,
	"dates":                  147.0000000000,
	"dill weed":              8.9000000000,
	"dressing":               245.4000000000,
	"egg noodles":            114.2000000000,
	"eggplant":               104.0000000000,
	"evaporated milk":        252.0000000000,
	"fennel":                 87.0000000000,
	"feta cheese":            150.0000000000,
	"flour":                  113.7000000000,
	"fruit":                  225.8000000000,
	"garlic":                 144.9000000000,
	"gelatin":                195.0000000000,
	"ginger":                 96.0000000000,
	"graham crackers":        84.0000000000,
	"green beans":            220.7000000000,
	"green chilies":          241.0000000000,
	"green onions":           71.0000000000,
	"green pepper":           130.0000000000,
	"green peppers":          238.5000000000,
	"ham":                    250.4000000000,
	"hamburger":              244.0000000000,
	"hazelnuts":              108.3000000000,
	"honey":                  59.1000000000,
	"ice":                    187.3000000000,
	"kale":                   117.0000000000,
	"kidney beans":           194.6000000000,
	"leeks":                  75.0000000000,
	"lemon":                  196.9000000000,
	"lemon juice":            244.0000000000,
	"lemons":                 212.0000000000,
	"lettuce":                49.2000000000,
	"lime":                   198.0000000000,
	"lime juice":             244.0000000000,
	"mango":                  251.0000000000,
	"maple syrup":            322.0000000000,
	"margarine":              225.7000000000,
	"marshmallows":           39.3000000000,
	"mayonnaise":             230.6000000000,
	"milk":                   220.3000000000,
	"molasses":               337.0000000000,
	"mushrooms":              116.8000000000,
	"mustard":                153.0000000000,
	"nuts":                   150.6000000000,
	"oatmeal":                60.5000000000,
	"oats":                   81.0000000000,
	"oil":                    207.3000000000,
	"olive oil":              216.0000000000,
	"onion":                  152.6000000000,
	"onions":                 176.5000000000,
	"orange":                 237.7000000000,
	"orange juice":           251.2000000000,
	"oranges":                176.0000000000,
	"parmesan":               126.7000000000,
	"parmesan cheese":        100.0000000000,
	"parsley":                32.8000000000,
	"pasta":                  116.0000000000,
	"peaches":                229.9000000000,
	"peanut butter":          154.4000000000,
	"peanut oil":             216.0000000000,
	"peanuts":                135.8000000000,
	"pears":                  190.6000000000,
	"peas":                   171.1000000000,
	"pecans":                 107.0000000000,
	"pepper":                 124.0000000000,
	"pie crust":              129.0000000000,
	"pine nuts":              135.0000000000,
	"pineapple":              235.3000000000,
	"pineapple juice":        250.0000000000,
	"pork":                   146.7000000000,
	"potato":                 191.7000000000,
	"potatoes":               180.4000000000,
	"pumpkin":                134.1000000000,
	"quinoa":                 177.5000000000,
	"raisins":                140.2000000000,
	"raspberries":            179.8000000000,
	"red pepper":             125.0000000000,
	"red wine vinegar":       239.0000000000,
	"rice":                   89.2000000000,
	"ricotta cheese":         247.0000000000,
	"salad":                  207.0000000000,
	"salad oil":              214.0000000000,
	"salmon":                 177.0000000000,
	"salsa":                  250.0000000000,
	"salt":                   292.0000000000,
	"sauce":                  251.4000000000,
	"sausage":                140.3000000000,
	"scallions":              100.0000000000,
	"semolina":               107.6000000000,
	"sesame oil":             218.0000000000,
	"sesame seeds":           144.0000000000,
	"shallots":               14.4000000000,
	"shortening":             207.3000000000,
	"shrimp":                 200.2000000000,
	"skim milk":              245.0000000000,
	"soda":                   57.5000000000,
	"sour cream":             237.2000000000,
	"soy sauce":              243.5000000000,
	"spaghetti":              151.4000000000,
	"spinach":                164.1000000000,
	"spray":                  247.0000000000,
	"sugar":                  205.8000000000,
	"sweet potatoes":         224.0000000000,
	"swiss cheese":           137.7000000000,
	"tofu":                   250.0000000000,
	"tomato":                 230.5000000000,
	"tomato juice":           241.2000000000,
	"tomato sauce":           217.1000000000,
	"tomato soup":            206.2000000000,
	"tomatoes":               177.6000000000,
	"tuna":                   146.0000000000,
	"turkey":                 206.6000000000,
	"vanilla":                175.9000000000,
	"vegetable broth":        225.3000000000,
	"vegetable oil":          211.5000000000,
	"vegetable shortening":   205.0000000000,
	"vinegar":                238.0000000000,
	"walnuts":                95.0000000000,
	"water":                  243.2000000000,
	"whipped cream":          70.0000000000,
	"white bread":            41.2000000000,
	"white rice":             172.9000000000,
	"worcestershire sauce":   275.0000000000,
	"yellow cornmeal":        143.3000000000,
	"yogurt":                 211.8000000000,
	"zucchini":               194.4000000000,
}
