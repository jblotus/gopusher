package gopusher

import (
    "testing"
)

func Test_Trigger(t *testing.T) {    
    channels := []string{"foo", "bar", "baz"}
    
    r, err := Trigger(channels);
    
    if (r != "this will return a pusher response eventually") {
        t.Errorf("expected `%s` to be `foo`", r);    
    }
    
    if (err != nil) {
        t.Errorf("some error happened");
    }
    
}

func Test_Trigger_ErrorsOnTooManyChannels(t *testing.T) {    
    channels := []string{
        "Shuffletag",
        "Vinte",
        "Jabbercube",
        "Edgeify",
        "Rhycero",
        "Buzzshare",
        "Edgeclub",
        "Npath",
        "Kazio",
        "Yadel",
        "Thoughtbeat",
        "Myworks",
        "Thoughtstorm",
        "Jayo",
        "Camido",
        "Photobug",
        "Twitterlist",
        "Layo",
        "Eidel",
        "Mynte",
        "Mynte",
        "Wikibox",
        "Devbug",
        "Agimba",
        "Zooxo",
        "Mynte",
        "Skipstorm",
        "Cogibox",
        "Quatz",
        "Ntag",
        "Photobug",
        "Realpoint",
        "Buzzster",
        "Lajo",
        "Rhynyx",
        "Eimbee",
        "Devpulse",
        "Rhybox",
        "Voomm",
        "Ntag",
        "Riffwire",
        "Izio",
        "Aimbu",
        "Feednation",
        "Skyble",
        "Realbuzz",
        "Kwinu",
        "Rhyloo",
        "Edgepulse",
        "Omba",
        "Livepath",
        "Yozio",
        "Skyble",
        "Edgewire",
        "Meevee",
        "Gabvine",
        "Zoombox",
        "Blogpad",
        "Voolith",
        "Avavee",
        "Topdrive",
        "Plambee",
        "Jaxbean",
        "Einti",
        "Ntags",
        "Skipfire",
        "Jabbertype",
        "Jaxworks",
        "Topicblab",
        "Zoonder",
        "Tazz",
        "Roodel",
        "Twinte",
        "Trilith",
        "Meeveo",
        "Topicblab",
        "Browsecat",
        "Dazzlesphere",
        "Mita",
        "Voomm",
        "Minyx",
        "Oodoo",
        "Tazzy",
        "Meevee",
        "Devcast",
        "Feedmix",
        "Quamba",
        "Pixonyx",
        "Edgewire",
        "Fiveclub",
        "Jabbersphere",
        "Linktype",
        "Twitterwire",
        "Browsedrive",
        "Mynte",
        "Skajo",
        "Browseblab",
        "Aimbo",
        "Gigashots",
        "Yakijo",
        "JooJoo",
    }
    
    r, err := Trigger(channels);
    
    if err == nil {
       t.Errorf("expected error due to too many channels, got `%s` instead", r);
    }
    
}