package lzw

import "testing"

func TestEncode(t *testing.T) {
	txt := `
	Lorem ipsum dolor sit amet, consectetur adipiscing elit. In id diam dui. Phasellus erat lorem, aliquet in nulla faucibus, consequat vestibulum sapien. Aenean sem enim, commodo quis dapibus a, placerat molestie lacus. Mauris elementum urna augue. Praesent at tellus vitae sem scelerisque interdum. Aenean id vulputate nisi, sit amet malesuada nisi. Ut consectetur eros in consectetur luctus. Ut blandit orci eu enim congue, vel dignissim ante venenatis. Nullam eu pharetra odio. Nam fringilla ante vel sem eleifend, luctus molestie ex sagittis. Aliquam elementum, arcu in dapibus cursus, tortor orci faucibus mauris, eget varius enim nulla auctor erat. Vivamus vehicula urna metus, sit amet feugiat leo rutrum in. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.

	In malesuada pharetra diam, eu fermentum ex egestas pharetra. Suspendisse ultrices ligula ac eros molestie, quis consectetur libero egestas. Vestibulum et faucibus dolor. Suspendisse sed purus dictum, imperdiet felis sed, dignissim sem. Etiam laoreet fringilla lorem condimentum consequat. Sed porttitor lorem pellentesque sem pretium, hendrerit convallis purus rhoncus. Nullam neque neque, eleifend nec eros ac, semper euismod arcu. Nulla interdum ex sagittis, tincidunt sem sed, fermentum ante. Proin in convallis enim, vitae consectetur arcu. Nulla non tellus id ligula mollis egestas. Duis a nibh ut odio rutrum molestie. Sed ultrices, ipsum ut ultricies dignissim, metus velit suscipit purus, eu vestibulum nibh nibh eu ipsum. Nam tristique et nulla at efficitur.
	
	Phasellus vel massa sit amet lacus lacinia ornare id sit amet nisi. Suspendisse potenti. Suspendisse augue libero, condimentum sodales felis vitae, vestibulum dapibus ex. Etiam lobortis, felis at pellentesque mattis, nisl augue mattis tortor, nec euismod dolor tellus quis arcu. Mauris rutrum porta nisi, sit amet iaculis velit porttitor ac. Proin mauris purus, vestibulum vitae risus sed, scelerisque congue lectus. Vestibulum aliquam turpis ac mi blandit, ut venenatis felis consectetur. In ex sapien, pharetra vitae mauris nec, convallis euismod libero. Sed tincidunt purus at tortor commodo volutpat.
	
	Donec tincidunt ac mi at pulvinar. Phasellus cursus dapibus sollicitudin. Nullam eleifend sagittis sem, quis auctor nisl facilisis vel. Quisque ornare lacinia posuere. Ut dictum sapien non risus convallis ultrices. Vivamus mollis id turpis vel molestie. Nam tempor odio vitae orci condimentum semper. Pellentesque sodales ante non turpis accumsan maximus. Aenean pellentesque lacinia tempor. Praesent sit amet varius enim. Duis auctor sollicitudin risus, et pulvinar magna tincidunt in. Quisque quis nibh ac nisi laoreet suscipit vehicula accumsan ligula. Ut dapibus vestibulum blandit. Nam consectetur orci justo, at gravida ante tristique ac. Praesent suscipit nunc vel pharetra pharetra.
	
	Donec porttitor viverra dui, eget auctor nisl tincidunt sit amet. Aliquam finibus, velit non finibus cursus, sem turpis ullamcorper diam, eu interdum lorem lectus tincidunt enim. Suspendisse ac purus ullamcorper, elementum magna ac, tincidunt nisl. Cras porta mattis purus, a condimentum turpis. Ut maximus lorem et lacus consequat, nec posuere nibh posuere. Maecenas sapien risus, tempus rhoncus magna id, pellentesque imperdiet elit. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc bibendum nisi quis odio molestie finibus. Duis et ante lacus. Sed blandit, metus a volutpat dictum, leo mi posuere ante, sit amet tempor dui lorem sit amet magna.
	`

	encoded := Encode(txt)

	if len(encoded) < 1 {
		t.Errorf("Encoding string of len: %d resulted in %d", len(txt), len(encoded))
	}
}
