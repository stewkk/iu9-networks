#ifndef CANVAS_H_
#define CANVAS_H_

#include <gtkmm/drawingarea.h>

namespace lab32 {
    class Canvas : public Gtk::DrawingArea {
        public:
            Canvas();
            virtual ~Canvas();
        private:
            void OnDraw(const Cairo::RefPtr<Cairo::Context>& cr, int width, int height);
    };
}

#endif // CANVAS_H_
